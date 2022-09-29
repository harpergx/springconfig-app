package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	myconfig    string `json:"config"`
	description string `json:"info.description"`
}

func Response(w http.ResponseWriter, data interface{}, status int, err error) {
	var res CommonResponse

	if status == http.StatusOK {
		res.Data = data
		res.Status = status
	} else {
		res.Status = status
		res.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func Controller(router *mux.Router) error {

	router.HandleFunc("/check/profile", func(w http.ResponseWriter, r *http.Request) {

		Response(w, "test", http.StatusOK, nil)

	}).Methods("GET")
	return nil
}
