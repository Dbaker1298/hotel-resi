build:
	@go build -o bin/api

run: build
	@./bin/api

docker-build:
	echo "Building docker image"
	@docker build -t dbaker1298/api .
seed:
	@go run scripts/seed.go

test:
	@go test -v ./...