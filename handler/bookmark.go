package handler

import (
	"encoding/json"
	"net/http"

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
