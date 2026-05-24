package internalSql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL(dbUser, dbPassword, dbHost, dbPort, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", dbUser, dbPassword, dbHost, dbPort, dbName, "America%2FSao_Paulo")

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("error connecting to database")
	}

	log.Println("database connected")

	return db, err
}
