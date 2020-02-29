# Go-book-store-api
Go Sample project to understand Mysql CRUD operation with best practises


A production ready sample Book store RESTful API with Go using **gorilla/mux** **uber/Zap** **lumberjack** with **Mysql** (A nice relational Database)

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
make run

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

```

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

## Todo

- [√] Support basic REST APIs.
- [ ] Support Authentication with user for securing the APIs.
- [√] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [√] Organize the code with packages
- [√] Add logging to application
- [ ] Add zipkin for tracing
- [ ] Make docs with GoDoc
- [ ] Setting up the swagger
- [ ] Dockerized the application
- [ ] Building a deployment process
