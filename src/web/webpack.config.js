import path from "path";
import WebpackMerge from "webpack-merge";
const __dirname = path.dirname("");

const baseConfig = {
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"],
    alias: {},
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx|ts|tsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "ts-loader",
        },
      },
      {
        test: /\.less$/,
        use: [
          {
            loader: "style-loader",
          },
          {
            loader: "css-loader",
          },
          {
            loader: "less-loader",
          },
        ],
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
      {
        test: /\.(jpg|png|jpeg|gif)$/,
        use: [
          {
            loader: "file-loader",
            options: {
              name: "[name]_[hash].[ext]",
              outputPath: "images/",
            },
          },
        ],
      },
    ],
  },
  plugins: [new CleanWebpackPlugin.CleanWebpackPlugin()],
  experiments: {
    outputModule: true,
  },
};

const devConfig = {
  mode: "development",
  devtool: "inline-source-map",
  entry: {
    index: "./src/index.tsx",
  },
  output: {
    filename: "[name].[chunkhash].js",
    path: path.resolve(__dirname, "dist"),
    library: {
      type: "module",
    },
  },
  devServer: {
    compress: true,
    host: "127.0.0.1",
    port: 4000,
    liveReload: true,
  },
  plugins: [
    new HtmlWebpackPlugin({
      //   template: path.resolve(__dirname, "resources/template.html"),
      filename: "index.html",
      chunks: ["index"],
    }),
  ],
};

const prodConfig = {};

const config = WebpackMerge.merge(baseConfig, devConfig);

export default config;
