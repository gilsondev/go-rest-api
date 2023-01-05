.PHONY: api-setup
api-setup:
	cd ./backend && go mod tidy

.PHONY: web-setup
web-setup:
	cd ./frontend && npm install

.PHONY: setup
setup: api-setup web-setup

.PHONY: api-serve
api-serve:
	cd ./backend && go run main.go

.PHONY: web-serve
web-serve:
	cd ./frontend && API_URL_DOMAIN=http://localhost:8000 npm start

.PHONY: serve
serve: api-serve web-serve
