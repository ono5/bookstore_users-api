package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Set Environgment from docker-compose
	MYSQL_USER     = "MYSQL_USER"
	MYSQL_PASSWORD = "MYSQL_PASSWORD"
	MYSQL_DATABASE = "MYSQL_DATABASE"
	MYSQL_PORT     = "MYSQL_PORT"
)

var (
	Client *sql.DB

	username = os.Getenv(MYSQL_USER)
	password = os.Getenv(MYSQL_PASSWORD)
	database = os.Getenv(MYSQL_DATABASE)
	port     = os.Getenv(MYSQL_PORT)
)

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		port,
		database,
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
