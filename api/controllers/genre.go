package controllers

import (
	hc "book-store-api/api/constants"
	"book-store-api/api/helper"
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Genre model struct
type Genre struct {
	genreRepo models.GenreRepo
}

func NewGenreHandler(db *driver.DB) *Genre {
	return &Genre{
		genreRepo: r.NewGenreRepo(db.SQL),
	}
}

/**
 * Function to create Genre
 * @param  {[http.Response, http.Request]}  [Http Request and Response]
 * @return {[Object]}   [Success or Error Object]
 */
func (g *Genre) CreateGenre(w http.ResponseWriter, r *http.Request) {

	genre := models.CreateGenre{}
	err := json.NewDecoder(r.Body).Decode(&genre)
	//validate the request
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	if genre.Name == "" {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, r.Body)
	}
	res, err := g.genreRepo.Create(r.Context(), &genre)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusCreated, res)
}

/**
 * function to update Genre
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *Genre) UpdateGenre(w http.ResponseWriter, r *http.Request) {

	req := models.Genre{}
	err := json.NewDecoder(r.Body).Decode(&req)
	//validate the request
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	if r.Body == nil || req.Id == nil || req.Name == nil || req.Status == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, r.Body)
		return
	}
	res, err := g.genreRepo.Update(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}

func (g *Genre) GetAll(w http.ResponseWriter, r *http.Request) {
	//get the query params
	var limit, offset int64

	if r.URL.Query().Get("limit") == "" {
		limit = hc.LIMIT
	} else {
		limit, _ = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	}

	if r.URL.Query().Get("offset") == "" {
		offset = hc.OFFSET
	} else {
		offset, _ = strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	}

	res, err := g.genreRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *Genre) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := g.genreRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (g *Genre) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := g.genreRepo.Delete(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}
