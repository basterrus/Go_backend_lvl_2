package render

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, v interface{}, status int) {
	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, v interface{}) {
	JSON(w, v, http.StatusOK)
}

func NotFoundError(w http.ResponseWriter) {
	response := struct {
		Status bool   `json:"status"`
		Error  string `json:"error"`
	}{false, "Not found."}

	JSON(w, response, http.StatusNotFound)
}

func BadRequestError(w http.ResponseWriter, err error) {
	response := struct {
		Status bool   `json:"status"`
		Error  string `json:"error"`
	}{false, err.Error()}

	JSON(w, response, http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter) {
	response := struct {
		Status bool   `json:"status"`
		Error  string `json:"error"`
	}{false, "Unexpected error occurred. Please try request later."}

	JSON(w, response, http.StatusInternalServerError)
}
