# go-dating-app

- config/ : contains JWT secret and DB configurations/credentials
- handler/ : route handler
- migrations/ : db (pgsql) migration script
- model/ : data structure mapping from table
- model/request : data structure mapping from http request payload
- repository/ : contains methods that handle business logic
- router/ : entry point definitions
- shared/ : only use for open db connection

## Setup & Installation

### Prequisites
- Go

### Database Migration

You may import/restore the migration script to PostgreSQL or using pgadmin.

### Run Project

```shell
go run main.go
```