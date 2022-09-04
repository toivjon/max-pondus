const path = require("path");
const HtmlWebPackPlugin = require("html-webpack-plugin");
const ReactRefreshWebpackPlugin = require("@pmmmwh/react-refresh-webpack-plugin")

module.exports = {
  mode: "development",
  entry: "./src/index.js",
  output: {
    filename: "bundle.[fullhash].js",
    path: path.resolve(__dirname, 'dist'),
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
    open: true,
    hot: true
  },
  devtool: "inline-source-map",
  plugins: [
    new HtmlWebPackPlugin({
      template: "public/index.html",
      favicon: "public/favicon.png"
    }),
    new ReactRefreshWebpackPlugin(),
  ],
};