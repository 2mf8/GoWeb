package data

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

var Db *sql.DB
var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "wr@#kequ5060", "the database password")
var iport *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "127.0.0.1", "the database server")
var user = flag.String("user", "sa", "the database user")

func init() {
	var err error

	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *iport)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=disable", *server, *user, *password, *iport)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	Db, err = sql.Open("mssql", connString)

	if err != nil {
		panic(err)
	}

	//defer Db.Close()
}
