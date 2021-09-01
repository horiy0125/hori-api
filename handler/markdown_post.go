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

	id, ok := vars["markdown_post_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	mid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	markdownPost, err := h.markdownPostUsecase.Show(mid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res := model.ShowMarkdownPostResponse(*markdownPost)

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
		markdownPost := model.ShowMarkdownPostResponse(m)
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

	return http.StatusOK, markdownPost, nil
}
