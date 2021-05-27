run:
	templ generate && go run cmd/main.go

hot-reload:
	# Uses https://github.com/cosmtrek/air
	air
