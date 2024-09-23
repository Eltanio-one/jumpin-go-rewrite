package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	database "github.com/Eltanio-one/jumpin-go-rewrite/src/db"
	"github.com/Eltanio-one/jumpin-go-rewrite/src/validate"
	"golang.org/x/crypto/bcrypt"
)

type registerGymRequest struct {
	GymName              string `json:"gymname"`
	Address              string `json:"address"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"confirmation"`
	Email                string `json:"email"`
}

func RegisterGym(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "registerGym", log.LstdFlags)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid HTTP method", http.StatusMethodNotAllowed, "only accepts POST method")
		http.Error(w, "incorrect HTTP method used", http.StatusMethodNotAllowed)
		return
	}

	// get request data
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to read request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestData registerGymRequest

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := database.InitialiseConnection()
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	dup, err := database.CheckDuplicate(db, "SELECT gymname FROM gyms WHERE gymname = $1", requestData.GymName)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if dup.Username != "" {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid query parameter", http.StatusBadRequest, "username already taken")
		http.Error(w, "username already taken", http.StatusBadRequest)
		return
	}

	err = validate.Email(requestData.Email)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Password(requestData.Password)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if requestData.Password != requestData.PasswordConfirmation {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, "password confirmation does not match given password")
		http.Error(w, "password confirmation does not match given password", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "failed to hash password", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// insert user into db
	err = database.Insert(db, "INSERT INTO gyms (username, email, hash, accountcreated, address) VALUES ($1, $2, $3, $4, NOW(), $5)",
		requestData.GymName,
		requestData.Email,
		hash,
		requestData.Address,
	)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Gym %s Registered Successfully", requestData.GymName)))
	w.WriteHeader(http.StatusOK)
}
