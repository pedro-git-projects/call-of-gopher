help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: starts development server, if no value is passed the default port is 8080
run:
	go run ./cmd/api/*.go

## opens the form in a browser
web:
	librewolf ./frontend/fetch/index.html

