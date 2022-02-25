package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO mitarbeiter(id, name, vorname) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1, "Florian", "Klima")
	if err != nil {
		log.Fatal(err)
	}

}
