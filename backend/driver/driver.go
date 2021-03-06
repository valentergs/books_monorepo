package driver

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

//ConnectDB will be an exported func
func ConnectDB() *sql.DB {

	var err error

	const (
		user     = "rodrigovalente"
		password = "Gustavo2012"
		// Quando rodar em Docker o HOST precisa ter o mesmo nome do container onde roda o Postgresql - nesse caso vai ficar postgres:5432
		host     = "postgres"
		port     = 5432
		dbname   = "livros"
	)

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db

}
