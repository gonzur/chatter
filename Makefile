lint:
	cd chatter_server && golangci-lint run && cd ..


update-pre-push:
	pre-commit install --hook-type pre-push