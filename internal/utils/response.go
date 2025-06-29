package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, message string, data interface{}){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"data": data,
	})
}