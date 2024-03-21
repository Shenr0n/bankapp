package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Shenr0n/bankapp/util"
	_ "github.com/lib/pq"
)

/*const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank_app?sslmode=disable"
)*/

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	// Establish connection to the db
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the db: ", err)
	}
	//Initialize test queries
	testQueries = New(testDB)

	os.Exit(m.Run())
}
