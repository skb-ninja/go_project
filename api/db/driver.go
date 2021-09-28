package db

import (
	"database/sql"
)

const db, err = sql.Open("mysql", "db_user:password@/test_db")

//defer db.Close()
