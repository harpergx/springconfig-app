package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
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
		tmp := viper.Get("config.name")
		Response(w, tmp, http.StatusOK, nil)

	}).Methods("GET")
	return nil
}
