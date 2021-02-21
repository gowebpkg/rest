package rest

import (
	"encoding/json"
	"net/http"
)

// AsJSON writes data as json to response writer
func AsJSON(rw http.ResponseWriter, data interface{}) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(jsonBody)
}
