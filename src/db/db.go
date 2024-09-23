package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Eltanio-one/jumpin-go-rewrite/src/config"

	_ "github.com/lib/pq"
)

type User struct {
	UserID         int
	Username       string
	Email          string
	Hash           string
	DateOfBirth    time.Time
	AccountCreated time.Time
}

var ErrDBQueryError = fmt.Errorf("error querying database")

func InitialiseConnection() (*sql.DB, error) {
	configFileName := `C:\Users\eltan\Programming Learning Projects\jumpin-go-rewrite-test\src\config\config.json`

	config, err := config.ReadConfigFile(configFileName)
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.Username, config.Database.Password, config.Database.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// test db is connected using ping
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func FetchUser(db *sql.DB, query string, args ...interface{}) (*User, error) {
	var user User
	err := db.QueryRow(query, args...).Scan(
		&user.UserID,
		&user.Username,
		&user.Email,
		&user.Hash,
		&user.DateOfBirth,
		&user.AccountCreated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, ErrUserDoesNotExist
		}
		return &User{}, err
	}
	return &user, nil
}

func CheckDuplicate(db *sql.DB, query string, args ...interface{}) (*User, error) {
	var user User
	err := db.QueryRow(query, args...).Scan(
		&user.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}
		return &User{}, err
	}
	return &user, nil
}

func Insert(db *sql.DB, query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

var ErrUserDoesNotExist = errors.New("user does not exist in database")
