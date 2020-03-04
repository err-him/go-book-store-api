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

type Author struct {
	authorRepo models.AuthorRepo
}

func NewAuthorHandler(db *driver.DB) *Author {
	return &Author{
		authorRepo: r.NewAuthorRepo(db.SQL),
	}
}

/**
 * [func description]
 * @param  {[type]} a [description]
 * @return {[type]}   [description]
 */
func (a *Author) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	req := models.AuthorCreateReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate the request
	if req.Name == nil || req.About.YearsActive == nil || req.About.Language == nil || req.About.Personal.Dob == nil || req.About.Personal.Info == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := a.authorRepo.Create(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusCreated, res)
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (a *Author) GetAll(w http.ResponseWriter, r *http.Request) {
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
	res, err := a.authorRepo.GetAll(r.Context(), limit, offset)
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
func (a *Author) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := a.authorRepo.GetOne(r.Context(), vars["id"])
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
func (a *Author) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := a.authorRepo.Delete(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (a *Author) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	req := models.AuthorCreateReq{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error(), err.Error())
		return
	}
	//validate Request
	if req.Id == nil || req.Name == nil || req.About.YearsActive == nil || req.About.Language == nil || req.About.Personal.Dob == nil || req.About.Personal.Info == nil || req.Status == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := a.authorRepo.Update(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusNoContent, res)
}
