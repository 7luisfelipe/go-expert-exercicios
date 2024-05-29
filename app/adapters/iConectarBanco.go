package adapters

import "database/sql"

type IConectarBanco interface {
	GetConn() *sql.DB
}
