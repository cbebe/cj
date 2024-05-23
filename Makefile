all:
	templ generate \
		--watch \
		--proxy="http://localhost:8080" \
		--cmd="go run ."

page:
	go run . build
