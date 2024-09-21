package database

import (
	"database/sql"
	"fmt"

	"github.com/Eltanio-one/jumpin-go-rewrite/src/config"

	_ "github.com/lib/pq"
)

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
