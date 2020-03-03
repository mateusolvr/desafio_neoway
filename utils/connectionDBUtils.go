package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tkanos/gonfig"
)

// Configuration : define the configuration struct
type Configuration struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

// ConnectDB : Connect to the Database
func ConnectDB() *sql.DB {

	configuration := Configuration{}
	err := gonfig.GetConf("./config.json", &configuration)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.Password, configuration.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}
