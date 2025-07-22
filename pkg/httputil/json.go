package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

const internalServerErrorJSON = `{"error":"internal server error"}`

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	encoded, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error: writeJSON: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerErrorJSON))
	} else {
		w.WriteHeader(status)
		w.Write(encoded)
	}
}
