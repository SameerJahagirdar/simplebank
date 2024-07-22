package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	_"github.com/lib/pq"
)

var testQueries *Queries

const dbDriver ="postgres"
const dbSource ="postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"

var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	testDB,err = sql.Open(dbDriver,dbSource)
	if(err!=nil){
		fmt.Print(err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
	
}
