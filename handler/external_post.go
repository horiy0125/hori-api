package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/horri1520/hori-api/model"
	"github.com/horri1520/hori-api/usecase"
	"github.com/horri1520/hori-api/util"
)

type ExternalPostHandler struct {
	externalPostUsecase *usecase.ExternalPostUsecase
}

func NewExternalPostHandler(externalPostUsecase *usecase.ExternalPostUsecase) *ExternalPostHandler {
	return &ExternalPostHandler{
		externalPostUsecase: externalPostUsecase,
	}
}

// GET /v1/external_posts/{external_post_id}
func (h *ExternalPostHandler) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["external_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	externalPost, err := h.externalPostUsecase.Show(rid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res := model.ExternalPost(*externalPost)

	return http.StatusOK, res, nil
}

// GET /v1/external_posts
func (h *ExternalPostHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	externalPosts, err := h.externalPostUsecase.Index()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var res model.IndexExternalPostResponse

	for _, e := range externalPosts {
		externalPost := model.ExternalPost(e)
		res.ExternalPosts = append(res.ExternalPosts, externalPost)
	}

	return http.StatusOK, res, nil
}

// POST /v1/external_posts
func (h *ExternalPostHandler) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var externalPost model.CreateExternalPostRequest
	if err := json.NewDecoder(r.Body).Decode(&externalPost); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	createdId, err := h.externalPostUsecase.Create(externalPost.Title, externalPost.Url, externalPost.ThumbnailUrl, externalPost.CategoryId, externalPost.PublishedAt)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, createdId, nil
}

// PUT /v1/external_posts/{external_post_id}
func (h *ExternalPostHandler) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["external_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var externalPost model.UpdateExternalPostRequest
	if err := json.NewDecoder(r.Body).Decode(&externalPost); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	if err := h.externalPostUsecase.Update(rid, externalPost.Title, externalPost.Url, externalPost.ThumbnailUrl, externalPost.CategoryId, externalPost.PublishedAt); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}

// DELETE /v1/external_posts/{external_post_id}
func (h *ExternalPostHandler) Destroy(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["external_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := h.externalPostUsecase.Destroy(rid); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}
