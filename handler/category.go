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

type CategoryHandler struct {
	categoryUsecase *usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase *usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		categoryUsecase: categoryUsecase,
	}
}

// GET /v1/categories/{category_id}
func (h *CategoryHandler) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["category_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	category, err := h.categoryUsecase.Show(rid)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	res := model.CategoryResponse(*category)

	return http.StatusOK, res, nil
}

// GET /v1/categories
func (h *CategoryHandler) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	categories, err := h.categoryUsecase.Index()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var res model.IndexCategoryResponse

	for _, c := range categories {
		category := model.CategoryResponse(c)
		res.Categories = append(res.Categories, category)
	}

	return http.StatusOK, res, nil
}

// POST /v1/categories
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	var category model.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	createdId, err := h.categoryUsecase.Create(category.Name)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusCreated, createdId, nil
}

// PUT /v1/categories/{category_id}
func (h *CategoryHandler) Update(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["category_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var category model.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "bad request body"}
	}

	if err := h.categoryUsecase.Update(rid, category.Name); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}

// DELETE /v1/categories/{category_id}
func (h *CategoryHandler) Destroy(_ http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	requestedId, ok := vars["category_id"]
	if !ok {
		return http.StatusBadRequest, nil, &util.HttpError{Message: "invalid path parameter"}
	}

	rid, err := strconv.ParseInt(requestedId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	if err := h.categoryUsecase.Destroy(rid); err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusNoContent, nil, err
}
