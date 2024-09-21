package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	database "github.com/Eltanio-one/jumpin-go-rewrite/src/db"
	"github.com/Eltanio-one/jumpin-go-rewrite/src/validate"
	"github.com/gorilla/sessions"

	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

var (
	Store *sessions.CookieStore
)

type loginRequest struct {
	Usermail string `json:"usermail"`
	Password string `json:"password"`
	Token    string `json:"recaptcha_token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "login", log.LstdFlags)

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

	var (
		requestData loginRequest
		user        *database.User
		query       string
	)

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	recaptcha, err := recaptcha.NewReCAPTCHA(os.Getenv("RECAPTCHA_SECRET"), recaptcha.V3, 100*time.Second)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to generate recaptcha instance", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = recaptcha.Verify(requestData.Token)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "recaptcha verification issue", http.StatusBadRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the data
	db, err := database.InitialiseConnection()
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, "unable to initiate connection to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if !validate.UsernameOrEmail(requestData.Usermail) {
		err = validate.Username(requestData.Usermail)
		if err != nil {
			logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		query = "SELECT userid, username, email, hash, dateofbirth, accountcreated FROM users WHERE username = $1"
	} else {
		err = validate.Email(requestData.Usermail)
		if err != nil {
			logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		query = "SELECT userid, username, email, hash, dateofbirth, accountcreated FROM users WHERE email = $1"
	}

	user, err = database.FetchRow(db, query, requestData.Usermail)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, "unable to initiate connection to database", http.StatusInternalServerError)
		return
	}

	err = validate.Password(requestData.Password)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate user hash against password
	if !validate.Hash(user.Hash, requestData.Password) {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid request parameter", http.StatusBadRequest, "password does not match database record for user")
		http.Error(w, "password does not match database record", http.StatusBadRequest)
		return
	}

	// set up session for user
	session, _ := Store.Get(r, "login")
	session.Values["authenticated"] = true
	session.Values["userID"] = user.UserID
	err = session.Save(r, w)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "session error", http.StatusInternalServerError, err.Error())
		http.Error(w, "unable to save session", http.StatusInternalServerError)
		return
	}

	// return 200
	w.WriteHeader(200)
}
