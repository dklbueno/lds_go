package utils

import (
	"encoding/json"
	. "lds/syslogLDS"
	"net/http"
	"strconv"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func StringToInt(str string) int {
	s, err := strconv.Atoi(str)
	if err != nil {
		GetInstanceLog().Alert(err.Error())
		panic(err)
	}
	return s
}

func StringToBool(str string) bool {
	b, err := strconv.ParseBool(str)
	if err != nil {
		GetInstanceLog().Alert(err.Error())
		panic(err)
	}
	return b
}
