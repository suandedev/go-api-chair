package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Chair struct {
	IdChair int
	Material string
	Design string
	Types string
}

func main(){
	// getAllChairs()
	getChairById(2)
	// updateChair(1, "Металл", "Крест", "Кухня")
	// deleteChair(1)
	// insertChair("Металл", "Крест", "Кухня")
}

// func insert chair
func insertChair(material string, design string, types string) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-chair")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO chair (material, design, type) VALUES (?, ?, ?)", material, design, types)
	if err != nil {
		log.Fatal(err)
	}
}

// func delete chair by id
func deleteChair(id int) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-chair")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM chair WHERE idChair = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}

// func update chair by id
func updateChair(id int, material string, design string, types string) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-chair")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE chair SET material = ?, design = ?, type = ? WHERE idChair = ?", material, design, types, id)
	if err != nil {
		log.Fatal(err)
	}
}

// func get chair by id
func getChairById(id int) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-chair")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var chair Chair
	err = db.QueryRow("SELECT * FROM chair WHERE idChair = ?", id).Scan(&chair.IdChair, &chair.Material, &chair.Design, &chair.Types)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", chair)
}

func getAllChairs() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go-api-chair")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM chair")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var chairs []Chair
	for rows.Next() {
		var chair Chair
		err := rows.Scan(&chair.IdChair, &chair.Material, &chair.Design, &chair.Types)
		if err != nil {
			log.Fatal(err)
		}
		chairs = append(chairs, chair)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, chair := range chairs {
		fmt.Printf("%+v\n", chair)
	}
}