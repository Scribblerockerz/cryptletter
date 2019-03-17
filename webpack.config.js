var path = require('path');
var webpack = require('webpack');

module.exports = {
    entry: {
        bundle: './assets/main.js',
        initial: './assets/initial.scss'
    },
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, 'public')
    },
    module: {
        rules: [{
            test: /\.css$/,
            use: ['style-loader', 'css-loader', 'postcss-loader']
        },{
            test: /\.scss$/,
            use: [
                'style-loader',
                'css-loader',
                'postcss-loader',
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
