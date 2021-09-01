package main

import (
	"encoding/json"
	"net/http"

	"github.com/horri1520/hori-api/util"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)

	// w.Header().Add("Access-Control-Allow-Headers", "*")
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		he := &util.HttpError{
			Message: err.Error(),
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(he)
		return
	}

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
