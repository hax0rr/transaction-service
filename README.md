# Transaction Service

Transaction Service helps to create user accounts and help them do transactions in their accounts.

## How to run

To start the application, clone this repo and run the below command
```shell
$ docker compose up
```
It already binds to local port 8080 for server and 5432 for postgres db, so you can test the APIs locally as well.

## API Documentation

To generate documentation, use

```shell
$ swag init -g ./app/routes.go -o ./docs
```

#### Local
- http://localhost:8000/swagger/index.html

