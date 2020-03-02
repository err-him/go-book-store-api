package controllers

import (
	hc "book-store-api/api/constants"
	"book-store-api/api/handler"
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"book-store-api/config/driver"
	"encoding/json"
	"net/http"
)

type Books struct {
	bookRepo models.BooksRepo
}

func NewBookHandler(db *driver.DB) *Books {
	return &Books{
		bookRepo: r.NewBookRepo(db.SQL),
	}
}

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
