COMPOSE = docker compose
EXEC = docker exec

test:
	go test ./...

up: test
	$(COMPOSE) up

down:
	$(COMPOSE) down

exec-api:
	$(EXEC) -it txn-svc-api /bin/bash

exec-db:
	$(EXEC) -it txn-svc-postgres /bin/bash -c "psql -U postgres -d transaction-service"



