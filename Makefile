help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: starts development server, if no value is passed the default port is 4000
run:
	go run ./cmd/api/*.go

## dev: starts the svelte development server on port 8080
dev:
	cd ./frontend/ && npm run dev

