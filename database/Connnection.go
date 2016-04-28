package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

func main() {
	var db *sql.DB
	var err error

	db, err = sql.Open("db2://10.118.21.180:50000/CTIE:currentSchema=IE_APP_ONLINE;")
	if err != nil {
	    return err
	}
}
