package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Error struct {
	Target string `json:"target_uuid"`
	Errors string `json:"aggregated_errors"`
}

func NewErrorsGatewayHandler(r *sync.Map) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		payload, _ := ioutil.ReadAll(req.Body)

		if targetName, found := vars["target_name"]; found {
			r.Store(targetName, string(payload))
			value, _ := r.Load(targetName)
			fmt.Printf("%s", value)
		} else {
			http.NotFound(w, req)
		}
	})
}

func NewErrorsListHandler(r *sync.Map) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		renderJSON(w, getErrors(r))
	})
}

func getErrors(r *sync.Map) []Error {
	errors := []Error{}
	r.Range(func(key, value interface{}) bool {
		errObject := Error{Target: key.(string), Errors: value.(string)}
		errors = append(errors, errObject)
		return true
	})
	return errors
}

func renderJSON(w http.ResponseWriter, value interface{}) error {
	valueJSON, err := json.Marshal(value)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(valueJSON))
	} else {
		http.Error(w, err.Error(), 500)
	}
	return err
}
