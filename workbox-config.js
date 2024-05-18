module.exports = {
	globDirectory: 'docs/',
	globPatterns: [
		'**/*.{js,html}'
	],
	swDest: 'docs/sw.js',
	ignoreURLParametersMatching: [
		/^utm_/,
		/^fbclid$/
	]
};
