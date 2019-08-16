# Go simple CRUD App
A RESTful API example for simple CRUD application with Go and PostgreSQL

## Docker run
```bash
# Build and Run
sudo docker-compose build
sudo docker-compose up -d
# API Endpoint : http://127.0.0.1:8080
```

## Manual run
```bash
# Creat new database in postgres. Run initdb.sql in new database
```

```bash
# Build
go get -d -v ./...
go build -v ./...
```
```bash
# Run
export APP_DB_HOST='your db host'
export APP_DB_NAME='your db name'
export APP_DB_USERNAME='your db username'
export APP_DB_PASSWORD='your db password'
./go-simple-crud
```

## API

#### /accounts
* `GET` : Get all accounts

#### /account
* `POST` : Create an account

#### /account/:id
* `GET` : Get an account
* `PUT` : Update an account
* `DELETE` : Delete an account

## CURL examples
```bash
# Create account
curl --header "Content-Type: application/json"   --request POST   --data '{"name":"testname","email":"test@test.com","api_token":"test_api"}'   http://localhost:8080/account
```
```bash
# Get account with id=1
curl --header "Content-Type: application/json"   --request GET  http://localhost:8080/account/1
```
```bash
# Update account with id=1
curl --header "Content-Type: application/json"   --request PUT   --data '{"name":"testname2","email":"test2@test.com","api_token":"test2_api"}'   http://localhost:8080/account/1
```
```bash
# Delete account with id=1
curl --header "Content-Type: application/json"   --request DELETE http://localhost:8080/account/1
```
```bash
# Get all accounts
curl --header "Content-Type: application/json"   --request GET  http://localhost:8080/accounts
```
