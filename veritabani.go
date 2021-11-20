package main

import (
	"database/sql"
	f "fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	Server   = "localhost"
	Port     = 1433
	User     = "sa"
	Password = 1453
	Db       = "Products"
)

func CheckDbConn() {

	var err error

	ConnString := f.Sprintf("server=%s;user id=%s;password=%d;port=%d;database=%s;",
		Server, User, Password, Port, Db)

	conn, err := sql.Open("sqlserver", ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	f.Printf("Connected!\n")
	defer conn.Close()
	option := 0
	f.Println("0.GET \n1.INSERT \n2.UPDATE \n3.DELETE")
	f.Scanln(&option)
	switch option {
	case 0:
		GetProducts(conn)
	case 1:
		 result,_:= CreateProduct(conn)

		f.Println(result)
	case 2:
		UpdateProduct(conn)
	case 3:
		DeleteProduct(conn)
	default:
		f.Println("Invalid operation request")
	}



}
