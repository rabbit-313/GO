package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct{
	Name string
	Age int
}

var Db *sql.DB

func main_() {

	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	// cmd := `CREATE TABLE IF NOT EXISTS persons(name STRING, age INT)`
	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// cmd := "UPDATE persons SET age = ? WHERE name = ?"
	// cmd := "SELECT * FROM persons where age = ?"
	cmd := "SELECT * FROM persons"
	rows, _ := Db.Query(cmd)
	defer rows.Close()

	var pp []Person

	for rows.Next(){
		var p Person
		err := rows.Scan(&p.Name, &p.Age)

		if err != nil{
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp{
		fmt.Println(p.Name, p.Age)
	}

	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	}else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

}
