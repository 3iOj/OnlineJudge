package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	util "github.com/3iOj/OnlineJudge/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBName))
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
