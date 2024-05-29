package databaseadapter

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Implements IConectarBanco
type MariaDbConectar struct {
}

func (m *MariaDbConectar) GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/ordersdb?parseTime=true&loc=America%2FSao_Paulo")
	//db, err := sql.Open("mysql", "root:root@tcp(container-mariadb:3306)/ordersdb?parseTime=true&loc=America%2FSao_Paulo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
