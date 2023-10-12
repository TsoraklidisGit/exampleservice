package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	if err = initSchema(db); err != nil {
		log.Fatal(err)
	}

	if err = populateData(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Initialization and data population completed successfully.")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}

func initSchema(db *sql.DB) error {
	//Create the "items" table
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT
		);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func populateData(db *sql.DB) error {
	// Insert sample data into the "items" table
	insertSQL := `
		INSERT INTO items (name) VALUES (?);
	`

	data := []string{"Item 1", "Item 2"}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, name := range data {
		_, err := stmt.Exec(name)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
