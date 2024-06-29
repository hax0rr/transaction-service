# Transaction Service

Transaction Service helps to create user accounts and do operations like deposit or withdrawal.

## How to run

To start the application, clone this repo and run the below docker command
```shell
docker compose up
```
Alternatively, you can also use below make commands
```shell
make up # to start the server
make exec-api # to exec into the api container
make exec-db # to exec into the postgres container
```

It binds to local port 8080 for server and 5432 for postgres db, so APIs and DB  can be tested locally as well.

## API Documentation
#### Local
- http://localhost:8080/swagger/index.html

## Database tables
- accounts
- transactions


