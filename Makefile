lint:
	cd chatter_server && golangci-lint run && cd ..


update-pre-push:
	pre-commit install --hook-type pre-push

dev-build-serv:
	cd chatter_server/cmd/development && go run ./main.go

dev-build-client:
	cd chatter-client && npm run start
