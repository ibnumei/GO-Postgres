## Requirement

1. PostgreSQL
2. The dependencies:

```go
  go get github.com/lib/pq
  go get github.com/joho/godotenv
```

- PQ is package that required to `connect Go to PostgreSQL database`
- Godotenv is package to read the `.env` file

## How to use

1. clone this project.
2. You have to make database `recordings` in the postgres. Run this to terminal:
   ```zsh
    psql -U postgres
    postgres=# CREATE database recordings;
    postgres=# exit
   ```
3. Then you can import this sql file to database with running this command:
   ```zsh
    psql -d recordings -f recordings.sql
   ```
4. Or you can run manually to create database, create table, and insert data with `DBeaver`, `PGAdmin`, and so on.
5. You can run:
   ```go
    go mod tidy
   ```
   to install the dependencies.
6. And finally, you can run with
   ```go
    go run main.go
   ```
7. You can just run the handler that do you want in `main.go`.
