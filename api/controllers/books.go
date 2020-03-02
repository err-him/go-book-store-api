package controllers

import (
	hc "book-store-api/api/constants"
	"book-store-api/api/handler"
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
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
	if req.Name == nil || req.ISBN == nil || req.Prices.OldPrice == nil || req.Prices.NewPrice == nil || req.Language == nil || req.PublisherId == nil || req.PublishedAt == nil || req.BookGenre.Id == nil || req.BookAuthor.Id == nil || req.Other.Quantity == nil || req.Other.Type == nil || req.Other.NumberPages == nil {
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
	if req.Id == 0 || req.Status == nil || req.Name == nil || req.ISBN == nil || req.Prices.OldPrice == nil || req.Prices.NewPrice == nil || req.Language == nil || req.PublisherId == nil || req.PublishedAt == nil || req.BookGenre.Id == nil || req.BookAuthor.Id == nil || req.Other.Quantity == nil || req.Other.Type == nil || req.Other.NumberPages == nil {
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
