var webpack = require('webpack');

module.exports = {
	entry: {
		'bundle': './src/js/main.js',
	},
	output: {
		filename: 'dist/js/[name].js'
	},
	watchOptions: {
		aggregateTimeout: 100
	},
	devtool: 'source-map',
	watch: true,
	module: {
		loaders: [
			{
				test: /\.js$/,
				exclude: /(node_modules|bower_components)/,
				use: {
					loader: 'babel-loader',
					options: {
						presets: [
						'es2015',
						 'react'
						 ],
						plugins: [
						'transform-object-rest-spread'
						]
					}
				}
			}, {
				test: /\.scss$/,
				loaders: ['style-loader', 'css-loader', 'autoprefixer-loader',  'sass-loader']
			}, {
				test: /\.css$/,
				loaders: [
				'style-loader', {
					loader: 'css-loader',
					options: {minimize: true}
				}],
			}, {
				test: /\.(woff2?|ttf|eot|svg|png)$/,
				loader: 'url-loader?limit=10000'
			}
		]
	}, plugins: [
		new webpack.ProvidePlugin({
			jQuery: 'jquery',
	        $: 'jquery',
	        jquery: 'jquery'
		})
	]
}