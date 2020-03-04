package controllers

import (
	hc "book-store-api/api/constants"
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
	"book-store-api/handler"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Books struct {
	bookRepo models.BooksRepo
}

func NewBookHandler(db *driver.DB) *Books {
	return &Books{
		bookRepo: r.NewBookRepo(db.SQL),
	}
}

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *Books) AddBook(w http.ResponseWriter, r *http.Request) {
	req := models.Books{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate the request
	if req.Name == nil || req.Image == nil || req.ISBN == nil || req.Prices.OldPrice == nil || req.Prices.NewPrice == nil || req.Language == nil || req.PublisherId == nil || req.PublishedAt == nil || req.BookGenre.Id == nil || req.BookAuthor.Id == nil || req.Other.Quantity == nil || req.Other.Type == nil || req.Other.NumberPages == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := b.bookRepo.Add(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusCreated, res)
}

/**
 * [func description]
 * @param  {[type]} b [description]
 * @return {[type]}   [description]
 */
func (b *Books) UpdateBook(w http.ResponseWriter, r *http.Request) {
	req := models.Books{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate the request
	if req.Id == 0 || req.Image == nil || req.Status == nil || req.Name == nil || req.ISBN == nil || req.Prices.OldPrice == nil || req.Prices.NewPrice == nil || req.Language == nil || req.PublisherId == nil || req.PublishedAt == nil || req.BookGenre.Id == nil || req.BookAuthor.Id == nil || req.Other.Quantity == nil || req.Other.Type == nil || req.Other.NumberPages == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := b.bookRepo.Update(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (b *Books) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err.Error())
		return
	}
	res, err := b.bookRepo.Delete(r.Context(), id)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (b *Books) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err.Error())
		return
	}
	res, err := b.bookRepo.GetBookDetailById(r.Context(), id)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (b *Books) GetAll(w http.ResponseWriter, r *http.Request) {
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
	res, err := b.bookRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (b *Books) SearchBook(w http.ResponseWriter, r *http.Request) {
	//get the query params
	var query string
	if r.URL.Query().Get("q") == "" {
		handler.HttpError(w, http.StatusInternalServerError, hc.INVALID_SEARCH_PARAM, nil)
		return
	} else {
		query = r.URL.Query().Get("q")
	}
	res, err := b.bookRepo.SearchBookByName(r.Context(), query)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
