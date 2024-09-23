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

// func hashPassword(password string) (string, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(hashedPassword), nil
// }

type registerRequest struct {
	Username             string `json:"username"`
	Email                string `jsom:"email"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"confirmation"`
	DateOfBirth          string `json:"dateOfBirth"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "register", log.LstdFlags)

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var requestData registerRequest

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check for duplicates
	db, err := database.InitialiseConnection()
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	dup, err := database.CheckDuplicate(db, "SELECT username FROM users WHERE username = $1", requestData.Username)
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

	// validate all data provided
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

	if requestData.DateOfBirth == "" {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusInternalServerError, "please provide your date of birth")
		http.Error(w, "please provide ysour date of birth", http.StatusInternalServerError)
		return
	}

	requestData.DateOfBirth = formatDate(requestData.DateOfBirth)

	_, err = validate.Date(requestData.DateOfBirth)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "failed to validate date of birth", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// insert user into db
	err = database.InsertUser(db, "INSERT INTO users (username, email, hash, dateofbirth, accountcreated, name) VALUES ($1, $2, $3, $4, NOW(), $5)",
		requestData.Username,
		requestData.Email,
		hash,
		requestData.DateOfBirth,
		requestData.Name,
	)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("User %s Registered Successfully", requestData.Username)))
}

func formatDate(date string) string {
	var d string
	for _, char := range date {
		charS := string(char)
		if charS != "T" {
			d += charS
			continue
		}
		break
	}
	return d
}
