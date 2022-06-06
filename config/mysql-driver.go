package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	// buat koneksi ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	db, err := sql.Open("mysql", "")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("success")
	}
	return db
}
