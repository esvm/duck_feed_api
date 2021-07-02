BASE_IMAGE=duck_feed_api

pipeline/docker/base:
	docker build -t $(BASE_IMAGE):latest .
	docker save --output $(BASE_IMAGE).docker $(BASE_IMAGE):latest

pipeline/lint:
	docker load --input ./$(BASE_IMAGE).docker
	docker-compose run $(BASE_IMAGE) golangci-lint run -v

pipeline/test:
	make infrastructure/raise
	make db/bootstrap
	docker load --input ./$(BASE_IMAGE).docker
	docker-compose run $(BASE_IMAGE)

infrastructure/raise:
	docker-compose up -d db

infrastructure/destroy:
	docker-compose down -d db

test:
	make infrastructure/raise
	docker-compose run $(BASE_IMAGE)

db/bootstrap:
	sleep 10
	docker-compose exec -T -e PGPASSWORD=password db psql -h localhost -U user -p 5432 -c "CREATE DATABASE duck_feeds"
	docker-compose exec -T -e PGPASSWORD=password db psql -h localhost -d duck_feeds -U user -p 5432 \
		-f /src/duck_feeds_service/store/postgres/schema.sql

start:
	docker-compose run --service-ports $(BASE_IMAGE) go run cmd/api_gateway/api_gateway.go

build_and_push:
	env GOOS=linux GOARCH=386 go build -a --ldflags="-s" -o cmd/api_gateway/bin/api_gateway cmd/api_gateway/api_gateway.go
	docker build -t duck-feed cmd/api_gateway/
	rm -rf cmd/api_gateway/bin/
	docker tag duck-feed:latest 349254485044.dkr.ecr.sa-east-1.amazonaws.com/duck-feed:latest
	docker push 349254485044.dkr.ecr.sa-east-1.amazonaws.com/duck-feed:latest
