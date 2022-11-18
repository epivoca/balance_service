package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/epivoca/balance_service/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cant't load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Service can't connect to the db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
