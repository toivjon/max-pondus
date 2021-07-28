const path = require("path");
const HtmlWebPackPlugin = require("html-webpack-plugin");

module.exports = {
  mode: "development",
  entry: "./src/index.js",
  output: {
    filename: "bundle.[fullhash].js",
  },
  module: {
    rules: [
      {
        test: /\.(js)$/,
        exclude: /node_modules/,
        loader: "babel-loader",
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      }
    ]
  },
  devServer: {
    host: "localhost",
    port: 3000,
    watchContentBase: true,
    open: true,
    hotOnly: true
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: "public/index.html",
    }),
  ],
};