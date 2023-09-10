package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var MaxxDB MaxxDataBase

type MaxxDataBase struct {
	Storage *DataStorage
}

type DataStorage struct {
	dbPath string
}

func NewDataStorage(dbPath string) *DataStorage {
	return &DataStorage{dbPath: dbPath}
}

func (ds *DataStorage) openDB() *sql.DB {
	db, err := sql.Open("sqlite3", ds.dbPath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (ds *DataStorage) Exec(query string, args ...interface{}) {
	db := ds.openDB()
	defer db.Close()
	_, err := db.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func (ds *DataStorage) InitializeTables() {
	tableQueries := []string{
		"CREATE TABLE IF NOT EXISTS todos (id INTEGER PRIMARY KEY, completed BOOL, task TEXT)",
		"CREATE TABLE IF NOT EXISTS notes (id INTEGER PRIMARY KEY, title TEXT, content TEXT)",
		"CREATE TABLE IF NOT EXISTS files (id INTEGER PRIMARY KEY, filename TEXT, data BLOB)",
		"CREATE TABLE IF NOT EXISTS pictures (id INTEGER PRIMARY KEY, filename TEXT, data BLOB)",
		"CREATE TABLE IF NOT EXISTS memory_dumps (id INTEGER PRIMARY KEY, description TEXT, data BLOB)",
		"CREATE TABLE IF NOT EXISTS hex_dumps (id INTEGER PRIMARY KEY, description TEXT, data TEXT)",
	}

	for _, query := range tableQueries {
		ds.Exec(query)
	}
}

func (ds *DataStorage) InsertData(tableName string, fields string, values ...interface{}) {
	placeholder := ""
	for i := 0; i < len(values); i++ {
		placeholder += "?"
		if i < len(values)-1 {
			placeholder += ", "
		}
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, fields, placeholder)
	ds.Exec(query, values...)
}

func (ds *DataStorage) DeleteData(tableName string, condition string, args ...interface{}) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s;", tableName, condition)
	ds.Exec(query, args...)
}

// ListItems generic function to list items from a table
func (ds *DataStorage) ListItems(tableName string, scanTarget ItemScanner) ([]ItemScanner, error) {
	var items []ItemScanner

	// Open the DB connection
	db := ds.openDB()
	defer db.Close()

	// Prepare SQL query
	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		// Make a new instance of the scanTarget type
		newItem := scanTarget.NewInstance()

		if err := newItem.ScanRow(rows); err != nil {
			return nil, err
		}
		items = append(items, newItem)
	}

	return items, nil
}

/*
// Delete note with a specific id
	storage.DeleteData("notes", "id = ?", 1)

	// Delete files with a specific filename
	storage.DeleteData("files", "filename = ?", "sample.txt")

	// Delete pictures with a specific filename
	storage.DeleteData("pictures", "filename = ?", "picture.jpg")

	// Delete memory dumps with a specific description
	storage.DeleteData("memory_dumps", "description = ?", "sample_memory_dump")

	// Delete hex dumps with a specific description
	storage.DeleteData("hex_dumps", "description = ?", "sample_hex_dump")

*/
