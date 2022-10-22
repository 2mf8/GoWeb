package main

import (
	"github.com/2mf8/GoWeb/router"
	"github.com/2mf8/GoWeb/data"
	_ "github.com/denisenkom/go-mssqldb"
)

func main(){
	defer data.Db.Close()
	r := router.InitRouter()
	r.Run(":8080")
}