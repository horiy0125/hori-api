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

type BookmarkHandler struct {
	bookmarkUsecase *usecase.BookmarkUsecase
}

func NewBookmarkHandler(bookmarkUsecase *usecase.BookmarkUsecase) *BookmarkHandler {
	return &BookmarkHandler{
		bookmarkUsecase: bookmarkUsecase,
	}
}

// GET /v1/bookmarks/{bookmark_id}
func (h *BookmarkHandler) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["bookmark_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	bookmark, err := h.bookmarkUsecase.Show(rid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res := model.ShowBookmarkResponse(*bookmark)

	return http.StatusOK, res, nil
}

// GET /v1/bookmarks
func (h *BookmarkHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	bookmarks, err := h.bookmarkUsecase.Index()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var res model.IndexBookmarkResponse

	for _, b := range bookmarks {
		bookmark := model.ShowBookmarkResponse(b)
		res.Bookmarks = append(res.Bookmarks, bookmark)
	}

	return http.StatusOK, res, nil
}

// POST /v1/bookmarks
func (h *BookmarkHandler) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var bookmark model.CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}

	}

	createdId, err := h.bookmarkUsecase.Create(bookmark.Url, bookmark.Description)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, createdId, nil
}
