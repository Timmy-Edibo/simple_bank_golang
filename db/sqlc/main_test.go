package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	// dbSource = "postgres://postgres:postgres@172.17.131.145:5432/simple_bank?sslmode=disable"

)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot cannect to the database: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}

// func TestMain(m *testing.M) {
// 	conn, err := sql.Open(dbDriver, dbSource)
// 	if err != nil {
// 		log.Fatal("Cannot cannect to the database: ", err)
// 	}

// 	testQueries = New(conn)
// 	os.Exit(m.Run())
// }
