var webpack = require("webpack");

module.exports = {
    entry: __dirname + "/src/index.js",
    output: {
        path: __dirname + "/dist/assets",
        filename: "bundle.js",
        publicPath: "assets"
    },
    devServer: {
        inline: true,
        contentBase: __dirname + "/dist",
        port: 8081
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                loader: ["babel-loader"]
            }, 
            {
                test: /\.json$/,
                loader: 'json-loader',
            }, 
            {
                test: /\.css$/,
                loader: 'style-loader!css-loader!autoprefixer-loader',
            }, 
            {
                test: /\.scss$/,
                loader: 'style-loader!css-loader!autoprefixer-loader!sass-loader',
            }, 
        ]
    }
}