package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	database "github.com/Eltanio-one/jumpin-go-rewrite/src/db"
	"github.com/Eltanio-one/jumpin-go-rewrite/src/validate"
)

type request struct {
	Usermail string `json:"usermail"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "login", log.LstdFlags)

	db, err := database.InitialiseConnection()
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, "unable to initiate connection to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if r.Method != http.MethodPost {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid HTTP method", http.StatusMethodNotAllowed, "only accepts POST method")
		http.Error(w, "incorrect HTTP method used", http.StatusMethodNotAllowed)
		return
	}

	// get the data out of the request that is sent from react.JS
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%w", "unable to read request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var requestData request

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%w", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate the data
	err = validate.Usermail(requestData.Usermail)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%w", "invalid request parameter", http.StatusBadRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// set the session for user

	// return 200
}
