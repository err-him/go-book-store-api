package controllers

import (
	"book-store-api/config/driver"
	hc "book-store-api/api/constants"
	"book-store-api/api/helper"
	"book-store-api/api/models"
	r "book-store-api/api/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Publishers struct {
	pubRepo models.PubsRepo
}

func NewPublisherHandler(db *driver.DB) *Publishers {
	return &Publishers{
		pubRepo: r.NewPubsRepo(db.SQL),
	}
}

/**
 * Method to create publication
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (p *Publishers) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	req := models.ReqPublish{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate Request
	if req.Name == nil || req.Meta == nil || req.Meta.FoundingDate == nil || req.Meta.Desc == nil {
		helper.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, r.Body)
		return
	}

	res, err := p.pubRepo.Create(r.Context(), &req)
	if err != nil {
		fmt.Println("err", err.Error())
		helper.HttpError(w, http.StatusInternalServerError, hc.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	helper.HttpResponse(w, http.StatusCreated, res)
}

func (p *Publishers) GetAll(w http.ResponseWriter, r *http.Request) {
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

	res, err := p.pubRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	helper.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (p *Publishers) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := p.pubRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	helper.HttpResponse(w, http.StatusOK, res)
}

/**
 * [func description]
 * @param  {[type]} g [description]
 * @return {[type]}   [description]
 */
func (p *Publishers) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := p.pubRepo.Delete(r.Context(), vars["id"])
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	helper.HttpResponse(w, http.StatusNoContent, res)
}

/**
 * [func description]
 * @param  {[type]} p [description]
 * @return {[type]}   [description]
 */
func (p *Publishers) UpdatePublisher(w http.ResponseWriter, r *http.Request) {
	req := models.UpdateReqPublish{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, err.Error(), err.Error())
		return
	}
	//validate Request
	if req.Id == nil || req.Name == nil || req.Meta == nil || req.Meta.FoundingDate == nil || req.Meta.Desc == nil || req.Status == nil {
		helper.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, r.Body)
		return
	}
	res, err := p.pubRepo.Update(r.Context(), &req)
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	helper.HttpResponse(w, http.StatusNoContent, res)
}
