package pgconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mikhailpachshenko"
	password = "1722"
	dbname   = "semi-project"
)

/* Подключение к БД  */
func ConnectionDatabase() *sql.DB {
	pSqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pSqlConn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
