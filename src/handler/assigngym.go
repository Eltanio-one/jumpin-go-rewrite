package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type requestData struct {
	Gym string `json:"gym"`
}

func AssignGym(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "assignGym", log.LstdFlags)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid HTTP method", http.StatusMethodNotAllowed, "only accepts POST method")
		http.Error(w, "incorrect HTTP method used", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to read request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData requestData

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
