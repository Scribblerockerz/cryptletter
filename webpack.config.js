var path = require('path');
var webpack = require('webpack');

module.exports = {
    entry: './assets/main.js',
    output: {
        filename: 'bundle.js',
        path: path.resolve(__dirname, 'public')
    },
    module: {
        rules: [{
            test: /\.css$/,
            use: [ 'style-loader', 'css-loader' ]
        },{
            test: /\.scss$/,
            use: [
                'style-loader',
                'css-loader',
                'sass-loader',
                'image-webpack-loader'
            ]
        },{
            test: /\.svg$/i,
            use: ['url-loader', 'image-webpack-loader']
        }]
    },
    plugins: [
      new webpack.optimize.UglifyJsPlugin()
    ]
};
