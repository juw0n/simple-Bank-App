package sqlcDB

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "user=root password=mypassword dbname=simple_Bank host=localhost port=5432 sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
