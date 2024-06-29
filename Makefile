COMPOSE = docker-compose

test:
	go test ./...

up: test
	$(COMPOSE) up -d

down:
	$(COMPOSE) down



