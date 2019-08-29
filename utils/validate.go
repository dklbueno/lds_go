package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"unsafe"
)

var StringJson string

func PaylodJson(w http.ResponseWriter, r *http.Request) bool {
	var log interface{}

	err := json.NewDecoder(r.Body).Decode(&log)
	switch {
	case err == io.EOF:
		http.Error(w, "Please send a request body", 400)
		return false
	case err != nil:
		http.Error(w, err.Error(), 400)
		return false
	}

	return true
}

func BodyToString(r *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	b := buf.Bytes()
	return *(*string)(unsafe.Pointer(&b))
}

func ValidateJson(w http.ResponseWriter, r *http.Request) bool {

	StringJson = BodyToString(r)

	if json.Valid([]byte(StringJson)) == false {
		http.Error(w, "Json Error", 400)
		return false
	}

	return true
}
