package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Club struct {
	name        string
	description string
}

func dbConnection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "yubaitao"
	dbPass := "yubaitao"
	dbName := "tcp(localhost:3306)/dbhw2"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getClubs() Club {
	db := dbConnection()
	rows, err := db.Query("SELECT cname, cdescription FROM Club WHERE cid = ?", 1)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var c Club
	for rows.Next() {
		err := rows.Scan(&c.name, &c.description)
		if err != nil {
			panic(err.Error())
		}
	}
	return c
}

func main() {
	fmt.Print("This is a test.\n")
	fmt.Print(getClubs().name + " - " + getClubs().description)
}
