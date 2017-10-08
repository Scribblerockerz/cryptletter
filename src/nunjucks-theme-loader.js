const fs = require('fs');
const path = require('path');
const nunjucks = require('nunjucks');

var NunjucksThemeLoader = nunjucks.FileSystemLoader.extend({
    getSource: function(name) {
        var fullpath = null;
        var paths = this.searchPaths;

        var useParent = !!name.match('parent:');
        var filename = name;

        if (useParent) {
            filename = name.replace('parent:', '');
        }

        for(var i=0; i<paths.length; i++) {
            var basePath = path.resolve(paths[i]);
            var p = path.resolve(paths[i], filename);

            // Only allow the current directory and anything
            // underneath it to be searched
            if(p.indexOf(basePath) === 0 && fs.existsSync(p)) {
                // Skip first iteration if use parent is called
                if (useParent && i === 0) continue;

                fullpath = p;
                break;
            }
        }

        if(!fullpath) {
            return null;
        }

        this.pathsToNames[fullpath] = name;

        return { src: fs.readFileSync(fullpath, 'utf-8'),
                 path: fullpath,
                 noCache: this.noCache };
    }
});

module.exports = NunjucksThemeLoader;
