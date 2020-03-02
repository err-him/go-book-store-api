package config

import (
	"book-store-api/api/controllers"
	"book-store-api/config/driver"
	"net/http"

	"github.com/gorilla/mux"
)

func handleAppRoutes(r *mux.Router, db *driver.DB) {

	genreHandler := controllers.NewGenreHandler(db)
	pubHandler := controllers.NewPublisherHandler(db)
	authorhandler := controllers.NewAuthorHandler(db)
	bookHandler := controllers.NewBookHandler(db)
	//handling API versioning
	v1 := r.PathPrefix("/api/v1").Subrouter()

	//genre routes
	v1.HandleFunc("/genre/create", genreHandler.CreateGenre).Methods(http.MethodPost)
	v1.HandleFunc("/genre/update", genreHandler.UpdateGenre).Methods(http.MethodPut)
	v1.HandleFunc("/genre/get/all", genreHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/genre/get/{id}", genreHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/genre/delete/{id}", genreHandler.Delete).Methods(http.MethodDelete)

	//Publishers routes
	v1.HandleFunc("/publishers/create", pubHandler.CreatePublisher).Methods(http.MethodPost)
	v1.HandleFunc("/publishers/get/all", pubHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/publishers/get/{id}", pubHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/publishers/delete/{id}", pubHandler.Delete).Methods(http.MethodDelete)
	v1.HandleFunc("/publishers/update", pubHandler.UpdatePublisher).Methods(http.MethodPut)

	//Author routes
	v1.HandleFunc("/author/create", authorhandler.CreateAuthor).Methods(http.MethodPost)
	v1.HandleFunc("/author/get/all", authorhandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/author/get/{id}", authorhandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/author/delete/{id}", authorhandler.Delete).Methods(http.MethodDelete)
	v1.HandleFunc("/author/update", authorhandler.UpdateAuthor).Methods(http.MethodPut)

	//Book Routes
	v1.HandleFunc("/book/add", bookHandler.AddBook).Methods(http.MethodPost)
}
