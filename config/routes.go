package config

import (
	"book-store-api/api/controllers"
	"book-store-api/config/driver"
	"book-store-api/middleware"
	mw "book-store-api/middleware"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func handleAppRoutes(r *mux.Router, db *driver.DB) {

	genreHandler := controllers.NewGenreHandler(db)
	pubHandler := controllers.NewPublisherHandler(db)
	authorhandler := controllers.NewAuthorHandler(db)
	bookHandler := controllers.NewBookHandler(db)
	userHandler := controllers.NewUserHandler(db)

	//api health check
	r.HandleFunc("/health", healthCheck).Methods(http.MethodGet)
	//handling API versioning
	v1 := r.PathPrefix("/api/v1").Subrouter()
	//genre routes
	v1.Handle("/genre/create", mw.JWTAuthMiddleware(http.HandlerFunc(genreHandler.CreateGenre))).Methods(http.MethodPost)
	v1.Handle("/genre/update", mw.JWTAuthMiddleware(http.HandlerFunc(genreHandler.UpdateGenre))).Methods(http.MethodPut)
	v1.HandleFunc("/genre/get/all", genreHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/genre/get/{id}", genreHandler.GetOne).Methods(http.MethodGet)
	v1.Handle("/genre/delete/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(genreHandler.Delete))).Methods(http.MethodDelete)

	//Publishers routes
	v1.Handle("/publishers/create", mw.JWTAuthMiddleware(http.HandlerFunc(pubHandler.CreatePublisher))).Methods(http.MethodPost)
	v1.HandleFunc("/publishers/get/all", pubHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/publishers/get/{id}", pubHandler.GetOne).Methods(http.MethodGet)
	v1.Handle("/publishers/delete/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(pubHandler.Delete))).Methods(http.MethodDelete)
	v1.Handle("/publishers/update", mw.JWTAuthMiddleware(http.HandlerFunc(pubHandler.UpdatePublisher))).Methods(http.MethodPut)

	//Author routes
	v1.Handle("/author/create", mw.JWTAuthMiddleware(http.HandlerFunc(authorhandler.CreateAuthor))).Methods(http.MethodPost)
	v1.HandleFunc("/author/get/all", authorhandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/author/get/{id}", authorhandler.GetOne).Methods(http.MethodGet)
	v1.Handle("/author/delete/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(authorhandler.Delete))).Methods(http.MethodDelete)
	v1.Handle("/author/update", mw.JWTAuthMiddleware(http.HandlerFunc(authorhandler.UpdateAuthor))).Methods(http.MethodPut)

	//Book Routes
	v1.Handle("/book/add", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.AddBook))).Methods(http.MethodGet)
	v1.Handle("/book/update", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.UpdateBook))).Methods(http.MethodGet)
	v1.Handle("/book/delete/{id}", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.DeleteBook))).Methods(http.MethodDelete)
	v1.Handle("/book/get/all", mw.JWTAuthMiddleware(http.HandlerFunc(bookHandler.GetAll))).Methods(http.MethodGet)
	//Jwt Authentication middleware
	v1.HandleFunc("/book/get/{id}", bookHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/book/search", bookHandler.SearchBook).Methods(http.MethodGet)

	//Users Routes
	v1.HandleFunc("/users/create", userHandler.CreateUser).Methods(http.MethodPost)
	v1.HandleFunc("/users/verify", userHandler.VerifyUser).Methods(http.MethodPost)

	//Api Key validation middleare for all routes
	v1.Use(middleware.ApiKeyMiddleware)

}

//methos to check api health status

func healthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Status: "up",
	})
}
