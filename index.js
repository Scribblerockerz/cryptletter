const path = require('path');
const fs = require('fs');
const express = require('express');
const bodyParser = require('body-parser');
const nunjucks = require('nunjucks');
const mysql = require('promise-mysql');
const randomstring = require("randomstring");
const md5 = require('md5');
const moment = require('moment');

const NunjucksThemeLoader = require('./src/nunjucks-theme-loader');


// Load configuration file
//-----------------------------------------------------------------------------

let configurationFile = './parameters.json';

if (process.argv[2]) {
  configurationFile = path.join(process.cwd(), process.argv[2]);
}

const configuration = require(configurationFile);



// mysql connection setup
//-----------------------------------------------------------------------------

let schemaFile = path.join(__dirname, 'schema.sql');
const schemaSql = fs.readFileSync(schemaFile).toString();

// Get database connection
const connect = function () {
  return mysql.createConnection(configuration.database)
    .then((conn) => {
      return conn;
    })
    .catch((error) => {
      throw error;
    });
}

// Run setup scripts, prepare databse
connect().then((conn) => {
  conn.query(schemaSql);
  conn.end();
})



// Cleanup old entries from the cli
//-----------------------------------------------------------------------------

const deleteOldEntriesSql = 'DELETE FROM messages WHERE active_until IS NOT NULL AND active_until < CURRENT_TIMESTAMP() OR created_at < DATE_SUB(NOW(), INTERVAL 1 MONTH)';

setTimeout(function () {
    connect().then((conn) => {
      // Perform actual query
      conn.query(deleteOldEntriesSql);
      conn.end();
    });
}, configuration.cleanupInterval);



// Express configuration
//-----------------------------------------------------------------------------

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'public')))

// Add custom assets directory
if (configuration.app.assets) {
    var assetsPath = path.join(process.cwd(), configuration.app.assets);
    if (fs.existsSync(assetsPath)) {
        app.use(express.static(assetsPath));
    }
}

// List of view directories
var themes = [path.join(__dirname, 'views')];

// Load theme directory from the configuration
if (configuration.app.templates) {
    var themePath = path.join(process.cwd(), configuration.app.templates);
    if (fs.existsSync(themePath)) {
        themes.unshift(themePath);
    }
}

var TemplateLoader = new NunjucksThemeLoader(themes, {
    watch: false,
    noCache: false
});

var env = new nunjucks.Environment(TemplateLoader, {
    autoescape: true
});

env.express(app);

app.set('view engine', 'njk');



// Helper methods and static data
//-----------------------------------------------------------------------------

const delays = {
  '15': '15min',
  '30': '30min',
  '60': '1h',
  '120': '2h',
  '1440': '24h',
}

const MODE_MINUTES = 'minutes';
const MODE_SECONDS = 'seconds';

const getHashedIp = (req, token) => {
  let ip = req.headers['x-forwarded-for'] || req.connection.remoteAddress;
  return md5(ip + token);
}



// Homepage
//-----------------------------------------------------------------------------

app.get('/', (req, res) => {
  res.render('index', { delays: delays});
});



// Styleguide
//-----------------------------------------------------------------------------

app.get('/styleguide', (req, res) => {
  res.render('styleguide');
});



// Create a new message from post data
//-----------------------------------------------------------------------------

app.post('/', (req, res) => {
  const message = req.body.message;
  const delay = delays[req.body.delay] && req.body.delay || 15;
  const token = randomstring.generate(64);
  const createdAt = new Date();

  const messageStruct = {
    text: message,
    token: token,
    mode: MODE_MINUTES,
    mode_value: delay,
    created_at: createdAt
  };

  connect().then((conn) => {
    conn.query('INSERT INTO messages SET ?', messageStruct);
    conn.end();
  }).then(() => {
    return res.json({ success: true, token: token });
  }).catch(() => {
    console.error(error);
    return res.json({ success: false });
  });
});



// Destroy message
//-----------------------------------------------------------------------------

app.delete('/:token/$', (req, res) => {

  const token = req.params.token;
  const clientIp = getHashedIp(req, token);

  connect().then((conn) => {
    conn.query('DELETE FROM messages WHERE token = ? AND accessable_ip = ?', [token, clientIp]);
    conn.end();

    return true;
  }).then((success) => {
    return res.json({ success: success });
  }).catch((error) => {
    console.error(error);
    return res.json({ success: false, error: 'Not found' });
  });
})



// Display a single message
//-----------------------------------------------------------------------------

app.get('/:token/$', (req, res) => {

  const token = req.params.token;
  const clientIp = getHashedIp(req, token);

  let connection;
  let activeUntil;

  connect().then((conn) => {
    connection = conn;

    // cleanup old messages
    conn.query(deleteOldEntriesSql);

    return conn.query('SELECT * FROM messages WHERE token = ? AND (active_until > CURRENT_TIMESTAMP() OR active_until IS NULL) LIMIT 1', token);
  }).then((rows) => {
    if (rows.length === 0) throw 'Message not found';

    return rows[0];
  }).then((message) => {
    let accessableIp = message.accessable_ip;
    let createdAt = message.created_at;
    let delay = message.mode_value;

    // can access this message
    if (accessableIp && accessableIp !== clientIp) {
      throw 'Access denied';
    }

    let unit = message.mode === MODE_MINUTES ? 'm' : 's';
    activeUntil = moment(createdAt).add(delay, unit).toDate();

    if (message.active_until === null) {
      connection.query('UPDATE messages SET ? WHERE token = ?', [{ active_until: activeUntil, accessable_ip: clientIp }, token]);
    }

    connection.end();

    return message;
  }).then((message) => {

    activeUntil = message.active_until || activeUntil;

    return res.render('show', {
      message: message.text,
      token: message.token,
      activeUntilTimestamp: (activeUntil*1),
      activeUntilDate: moment(activeUntil).format('MMMM Do YYYY, h:mm:ss a'),
      timeRemaining: moment(activeUntil).fromNow()
    });

  }).catch((error) => {
    connection && connection.end();
    console.error(error);
    return res.render('404');
  });
});



// 404 nothing found
//-----------------------------------------------------------------------------

app.use(function(req, res, next){
  res.status(404);

  res.format({
    html: function () {
      res.render('404');
    },
    json: function () {
      res.json({ error: 'Not found' });
    },
    default: function () {
      res.type('txt').send('Not found');
    }
  })
});



// Start server
//-----------------------------------------------------------------------------

app.listen(configuration.server.port, () => {
  console.log(`Starting server on http://0.0.0.0:${configuration.server.port}`);
})
