package httpResp

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJson(w, code, map[string]string{"error": message})

}

//if you use interface datatype, parameter will be flexible. It can be sent in map, struct

func RespondWithJson(w http.ResponseWriter, code int, data interface{}) {
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)

}
