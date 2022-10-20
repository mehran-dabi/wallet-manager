# Wallet Manager

This project is a simple wallet manager where you can:
- create a wallet
- add funds to the wallet
- subtract funds from the wallet
- get wallets balance

## Structure
- The Domain-Driven Design is used to implement this project.
- The project uses `MYSQL` as a database because the scale of data is small.
It also has the ACID properties, which prevents the user to spend the same fund more than once.
- Also, `gin-gonic` is used to serve an HTTP server.

### packages
- __Config Package:__ This package contains the database and service configurations.
- __Domain Package:__ This package contains the services' entities, repositories, and logic.
- __Infrastructure Package:__ In this package the database connection is initialized.
  The following SQL script creates the `users` table:
```sql
CREATE TABLE IF NOT EXISTS wallets (
    id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    balance INT(32) NOT NULL DEFAULT 0,
    user_id INT(32) NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
    );

```
- __Mocks Package:__ This package mocks the behaviour of the interfaces for unit testing.
- __.golangci.yml:__ This file is the configuration for golangci lint.


## How to Run

First, we have to get the database up and running. So we use the docker-compose file provided.
```shell
make docker.up
```

This will create a local MySQL database on port `3306`.
You don't need to worry about the tables. After running the program, the migrations are automatically done while initializing the database connection.

Then we get to run the main program.
```shell
make run
```

This will start a HTTP server on port `:8080`

## Test Coverage
By running the following command you can run all the tests in the project:
```shell
make test
```

Unit testing is used here. I could achieve the following coverages:
- `repository`: 86.0%.
- `service`: 84.2%.
- `utils`: 50.0%.

## How to Use
A postman collection is prepared with detailed examples to see how to user the APIs.
It can be found [here](https://www.getpostman.com/collections/d333878cacfc5bd7da2f)

