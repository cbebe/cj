all:
	templ generate \
		--watch \
		--proxy="http://localhost:8080" \
		--cmd="go run ."

page:
	go run . build

docs/BabelStoneHan.woff2:
	wget -O $@ https://www.babelstone.co.uk/Fonts/WOFF/BabelStoneHan.woff2
