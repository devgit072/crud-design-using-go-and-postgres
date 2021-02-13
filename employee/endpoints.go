package employee

import (
	"database/sql"
)
import _ "github.com/lib/pq"

const (
	hostname = "10.14.31.164"
	port     = 5432
	username = "postgres"
	password = "yoyo99"
	dbName   = "postgres"
)

const POSTGRES_DRIVER_NAME = "postgres"

var dbConn *sql.DB

type DBMethods interface {
	createConnection()
	closeConection()
}

type Database struct {
}
