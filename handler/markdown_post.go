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

type MarkdownPostHandler struct {
	markdownPostUsecase *usecase.MarkdownPostUsecase
}

func NewMarkdownPostHandler(markdownPostUsecase *usecase.MarkdownPostUsecase) *MarkdownPostHandler {
	return &MarkdownPostHandler{
		markdownPostUsecase: markdownPostUsecase,
	}
}

// GET /v1/markdown_posts/{markdown_post_id}
func (h *MarkdownPostHandler) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["markdown_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	markdownPost, err := h.markdownPostUsecase.Show(rid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res := model.MarkdownPost(*markdownPost)

	return http.StatusOK, res, nil
}

// GET /v1/markdown_posts
func (h *MarkdownPostHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	markdownPosts, err := h.markdownPostUsecase.Index()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var res model.IndexMarkdownPostResponse

	for _, m := range markdownPosts {
		markdownPost := model.MarkdownPost(m)
		res.MarkdownPosts = append(res.MarkdownPosts, markdownPost)
	}

	return http.StatusOK, res, nil
}

// POST /v1/markdown_posts
func (h *MarkdownPostHandler) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var markdownPost model.CreateMarkdownPostRequest
	if err := json.NewDecoder(r.Body).Decode(&markdownPost); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	createdId, err := h.markdownPostUsecase.Create(markdownPost.Title, markdownPost.Body, markdownPost.CategoryId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, createdId, nil
}

// PUT /v1/markdown_posts/{markdown_post_id}
func (h *MarkdownPostHandler) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["markdown_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var markdownPost model.UpdateMarkdownPostRequest
	if err := json.NewDecoder(r.Body).Decode(&markdownPost); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	if err := h.markdownPostUsecase.Update(rid, markdownPost.Title, markdownPost.Body, markdownPost.CategoryId); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}

// DELETE /v1/markdown_posts/{markdown_post_id}
func (h *MarkdownPostHandler) Destroy(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["markdown_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := h.markdownPostUsecase.Destroy(rid); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}
