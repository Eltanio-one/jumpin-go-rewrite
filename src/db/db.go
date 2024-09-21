package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Eltanio-one/jumpin-go-rewrite/src/config"

	_ "github.com/lib/pq"
)

type User struct {
	UserID         int
	Usermail       string
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

func FetchRow(db *sql.DB, query string, args ...interface{}) (*User, error) {
	var user User
	err := db.QueryRow(query, args...).Scan(
		&user.UserID,
		&user.Usermail,
		&user.Hash,
		&user.DateOfBirth,
		&user.AccountCreated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, fmt.Errorf("user does not exist in database")
		}
		return &User{}, err
	}
	return &user, nil
}
