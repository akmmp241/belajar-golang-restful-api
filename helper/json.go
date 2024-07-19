package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, v any) {
	err := json.NewDecoder(r.Body).Decode(&v)
	PanicIfErr(err)
}

func WriteToResponseBody(w http.ResponseWriter, v any) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&v)
	PanicIfErr(err)
}
