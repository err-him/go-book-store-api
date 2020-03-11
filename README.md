# Go-book-store-api
Go Sample project to understand Mysql CRUD operation with best practises

A production ready sample Book store RESTful API with Go using **gorilla/mux** **uber/Zap** **lumberjack** with **Mysql** (A nice relational Database), JWT Authentication. This Project contains a golang implementation of Swagger 2.0 (aka OpenAPI 2.0): it knows how to serialize and deserialize swagger specifications.

## Installation & Run
```bash
# Download this project
git clone github.com/err-him/go-book-store-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [db/evv.local.json](https://github.com/err-him/go-book-store-api/blob/master/config/properties/db/env.local.json)

```
{
  "host"   :  "127.0.0.1",
	"port"   :  "3306",
	"uname"  :  "root",
	"dbname" :  "book_store",
	"pass"   :  "root"
}
```

## Run via Docker

### Requirements

**Step 1:** Create the Docker image according to [Dockerfile](Dockerfile).
Ensure docker-compose is installed on your build system.
For details on how to do this, see: https://docs.docker.com/compose/install/

```shell
# This may take a few minutes.
$ docker-compose up -d
```
This will create two containers in background for Go and Mysql respectively

**Step 2:**.
The above process will create two separate container and DB for application, to auto populate some data in DB run

```shell

$ cat db.sql | docker exec -i full_db_mysql /usr/bin/mysql -u root --password=root book_store

```

**Step 3:** Open another terminal and access the example API endpoint.

```shell
$ curl http://localhost:9002/health
{"status": "up" }
```
**Important Note:** While setting up the docker change the DB host from ``127.0.0.1`` to ``book-store-mysql`` because while creating the docker Image we are proving host name ``book-store-mysql`` else use ``127.0.0.1`` if running using go run or make run


## Routing

**gorilla/mux**  is being used to setup for routing. It provides some powerful feature like grouping/middleware/handler etc.
``` Routing

  v1 := r.PathPrefix("/api/v1").Subrouter()

	v1.HandleFunc("/genre/create", genreHandler.CreateGenre).Methods(http.MethodPost)
	v1.HandleFunc("/genre/update", genreHandler.UpdateGenre).Methods(http.MethodPut)
	v1.HandleFunc("/genre/get/all", genreHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/genre/get/{id}", genreHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/genre/delete/{id}", genreHandler.Delete).Methods(http.MethodDelete)

```

## DB

Mysql is being used as database
**database/sql** and **github.com/go-sql-driver/mysql** module to create and manage database connection



## Build

```bash
# Build and Run
cd go-book-store-api
make build
```

## Run

```bash
# Build and Run
cd go-book-store-api
make APP_ENV="local" run

# API Endpoint : http://127.0.0.1:9002
```

## Structure
```
main.go -> Entry point of application
config -> folder to store all connection and routing related logic

  config/app.go -> app run/db connection and routing Initialize
  config/routes.go -> app routing defined here
  config/driver/db.go -> mysql connection established Here
  config/properties -> all application properties in JSON form stored Here, to read config file **https://github.com/err-him/gonf** package used

logger -> Folder contains application logging logic
    logger/zap.go -> This contains the logger implementation logic. To implement **uber/zap** logger with **lumberjack** a log rolling package **https://github.com/err-him/gozap** package is being used

api  -> Api package is used to receive an incoming request, validate the request for any bad input parameters. Generate a proper     response after running our business logic.

    api/constants  ->   contains all application related constants like http etc
    api/controllers ->  Contains handler functions for particular route to be called when an api is called.
    api/models ->   database tables to be used as models struct and interface provided for the repositories
    api/handler ->      basically contains the helper functions used in returning api responses, HTTP status codes, default messages etc.
    api/repositories ->   repository package is a wrapper on database and cache, so no other package can directly access the database. This package handle all create, update, fetch and delete operation on database tables or cache.
    api/utils ->    contains all application utility function.

db.sql -> file contains db schema and data information, dumped data from mysql
          command for dump database => `mysqldump -u root -p book_store > /Users/himanshu/go-learning/book-store-api/db.sql`

swagger.yml -> swagger API documentation configuration          

```
## OpenAPI & Swagger - API DOCUMENTATION
### ***Introduction to OpenAPI & Swagger Open Source Tools***

## OpenAPI & Swagger

### OpenAPI

**OpenAPI Specification** (formerly Swagger Specification) is an API description format for REST APIs. An OpenAPI file allows you to describe your entire API, including:

* Available endpoints (```/users```) and operations on each endpoint (```GET /users```, ```POST /users```)
* Operation parameters Input and output for each operation
* Authentication methods
* Contact information, license, terms of use and other information.

API specifications can be written in YAML or JSON. The format is easy to learn and readable to both humans and machines. The complete OpenAPI Specification can be found on GitHub:
[OpenAPI 2.0 Specification](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md),
[OpenAPI 3.0 Specification](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.0.md)

### Swagger

Swagger is a set of open-source tools built around the OpenAPI Specification that can help you design, build, document and consume REST APIs. The major Swagger tools include:
Swagger allows you to describe the structure of your APIs so that machines can read them. The ability of APIs to describe their own structure is the root of all awesomeness in Swagger.

* [Swagger Editor](http://editor.swagger.io/?_ga=2.27098621.139862542.1529283950-1958724428.1521772135) – browser-based editor where you can write OpenAPI specs.
* [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) – generates server stubs and client libraries from an OpenAPI spec.
* [Swagger UI](https://swagger.io/swagger-ui/) – renders OpenAPI specs as interactive API documentation.

## Introduction to OpenAPI Specification

### **Basic Structure**
Swagger can be written in JSON or YAML. In this guide, we only use YAML examples, but JSON works equally well. A sample Swagger specification written in YAML looks like:

```yaml
swagger: "2.0"
info:
  title: Sample API
  description: API description in Markdown.
  version: 1.0.0
host: api.example.com
basePath: /v1
schemes:
  - https
paths:
  /users:
    get:
      summary: Returns a list of users.
      description: Optional extended description in Markdown.
      produces:
        - application/json
      responses:
        200:
          description: OK
```



## Quick Start Swagger

 This Project contains a golang implementation of Swagger 2.0 (aka OpenAPI 2.0): it knows how to serialize and deserialize swagger specifications.

 `go-swagger` brings to the go community a complete suite of fully-featured, high-performance, API components to work with a Swagger API

**Installation Go Swagger**
To install Go SWAGGER in Mac, type following command

```
brew tap go-swagger/go-swagger
brew install go-swagger
```
Once Installation complete go to project repo and generate swagger documentation by following command

` swagger serve -F=swagger ./swagger.yml `

Also created nice Makefile for the same, to run via make enter:
`make serve-swagger`

** Swagger UI**
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-04%20at%2012.15.37%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.14%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.37%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.45%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.58%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.58%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.43.07%20PM.png)











## Quick Start Swagger

 This Project contains a golang implementation of Swagger 2.0 (aka OpenAPI 2.0): it knows how to serialize and deserialize swagger specifications.

 `go-swagger` brings to the go community a complete suite of fully-featured, high-performance, API components to work with a Swagger API

**Installation Go Swagger**
To install Go SWAGGER in Mac, type following command

```
brew tap go-swagger/go-swagger
brew install go-swagger
```
Once Installation complete go to project repo and generate swagger documentation by following command

` swagger serve -F=swagger ./swagger.yml `

Also created nice Makefile for the same, to run via make enter:
`make serve-swagger`

** Swagger UI**
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-04%20at%2012.15.37%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.14%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.37%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.45%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.58%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.42.58%20PM.png)
![Swagger UI](https://github.com/err-him/go-book-store-api/blob/develop/assets/swagger/Screenshot%202020-03-03%20at%206.43.07%20PM.png)









## API

#### /genre
* **/create** `POST` : Create Genre
* **/update** `PUT` : update  Genre
* **/get/all** `GET` : Get all active Genre
* **get/{:id}** `GET` : Get One genre - specified by Id or Genre slug
* **delete/{:id}** `DELETE` : delete genre - Soft delete by specified id or slug

#### /publishers
* **/create** `POST` : Create Publisher
* **/update** `PUT` : update  Publisher
* **/get/all** `GET` : Get all active Publisher
* **get/{:id}** `GET` : Get One Publisher - specified by Id or Genre slug
* **delete/{:id}** `DELETE` : delete Publisher - Soft delete by specified id or slug

#### /author
* **/create** `POST` : Create Author
* **/update** `PUT` : update  Author
* **/get/all** `GET` : Get all active Author
* **get/{:id}** `GET` : Get One Author - specified by Id or Genre slug
* **delete/{:id}** `DELETE` : delete Author - Soft delete by specified id or slug


#### /books
* **/create** `POST` : Create Book
* **/update** `PUT` : update  Book
* **/get/all** `GET` : Get all Books
* **get/{:id}** `GET` : Get One Book - specified by Id
* **delete/{:id}** `DELETE` : delete Book - Soft delete by Id
* **search?q={:query}** `GET` : Search book - by its Name

#### /users
* **/create** `POST` : Create User
* **/verify** `PUT` : Verify  Book


## Todo

- [√] Support basic REST APIs.
- [√] Support Authentication with user for securing the APIs.
- [√] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [√] Organize the code with packages
- [√] Add logging to application
- [ ] Add zipkin for tracing
- [√] Api documentation with Swagger
- [ ] Dockerized the application
- [ ] Building a deployment process

## Issues Management

Feel free to open an issue if you come across any bugs or
if you'd like to request a new feature.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b new-feature`)
3. Commit your changes (`git commit -am 'Some cool changes'`)
4. Push to the branch (`git push origin new-feature`)
5. Create new Pull Request
