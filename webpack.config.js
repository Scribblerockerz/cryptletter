var path = require('path');
var webpack = require('webpack');

module.exports = {
    mode: process.env['NODE_ENV'] === 'production' ? 'production' : 'development',
    entry: {
        bundle: './theme/assets/main.js',
        initial: './theme/assets/initial.scss',
    },
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, 'public'),
    },
    optimization: {
        minimize: process.env['NODE_ENV'] === 'production',
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader', 'postcss-loader'],
            },
            {
                test: /\.scss$/,
                use: ['style-loader', 'css-loader', 'postcss-loader', 'sass-loader', 'image-webpack-loader'],
            },
            {
                test: /\.svg$/i,
                use: ['url-loader', 'image-webpack-loader'],
            },
            {
                test: /\.m?js$/,
                exclude: /(node_modules)/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env'],
                    },
                },
            },
        ],
    },
};
