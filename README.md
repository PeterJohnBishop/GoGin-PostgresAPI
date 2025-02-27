# symmetrical-fishstick-go

A Go server using Gin and JWT to deploy a secured API to perform CRUD operations in a Postgres database.

## dependencies 

github.com/joho/godotenv
github.com/gin-gonic/gin
github.com/lib/pq
github.com/golang-jwt/jwt/v5
golang.org/x/crypto/bcrypt

## .env setup

HTTP_PORT=
GIN_PORT=
PSQL_HOST=
PSQL_PORT=
PSQL_USER=
PSQL_PASSWORD=
PSQL_DBNAME=
JWT_SECRET

