package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:a123456@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries = *db.Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("connect sql fail ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
