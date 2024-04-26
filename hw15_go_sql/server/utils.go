package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// M is a generic map.
type M map[string]interface{}

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(err)
	}
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	errorResponse(w, http.StatusInternalServerError, "internal error")
}

func errorResponse(w http.ResponseWriter, code int, errs interface{}) {
	writeJSON(w, code, M{"errors": errs})
}
