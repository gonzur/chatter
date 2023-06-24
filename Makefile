# Linter recipes
lint-server:
	cd chatter_server;\
	gofmt -s -w .&&\
	golangci-lint run

lint-client:
	cd chatter-client;\
	npm run lint

lint-all: lint-server lint-client
#

# Development recipes
run-server-dev:
	cd chatter_server/cmd/development;\
	go run ./main.go

run-client-dev:
	cd chatter-client;\
	npm run dev
#

# First run recipes
update-pre-push:
	pre-commit install --hook-type pre-push

setup-client:
	cd chatter-client;\
	npm i

setup-server:
	cd chatter_server;\
	go mod download

setup: setup-client setup-server update-pre-push lint-all
	git config --global rebase.autosquash true
#