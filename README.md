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
 After running `docker compose up` OR `make up`, the API documentation can be accessed via below link
- http://localhost:8080/swagger/index.html

## Database tables
Below tables can be accessed after running `make up` and `make exec-db` in your terminals - 
- accounts (example query: `select * from accounts limit 1;`)
- transactions (example query: `select * from transactions limit 1;`)


