package database

import (
	"database/sql"
	"ecom/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(conn)
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
