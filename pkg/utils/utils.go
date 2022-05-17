package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func WriteHeader(w http.ResponseWriter, statusCode int, message ...[]byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if len(message) > 0 {
		w.Write(message[0])
	}
}
