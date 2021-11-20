package main

import (
	"database/sql"
	f "fmt"
	"strings"
)

func GetProducts(db *sql.DB) (int, error) {

	getProduct_sql := "select * from Products"

	rows, err := db.Query(getProduct_sql)
	if err != nil {
		f.Println("Error reading records: ", err.Error())
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var name string
		var price float64
		var id int
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			f.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		f.Printf("ID: %d, Name: %s, Price: %f\n", id, name, price)
		count++
	}
	return count, nil
}

func CreateProduct(db *sql.DB)(int64,error){

	var name string
	f.Print("Please enter your product name: ")
	f.Scanln(&name)

	var price float64
	f.Print("Please enter your product's price: ")
	f.Scanln(&price)

	insertProduct_sql := f.Sprintf("INSERT INTO Products (name,price) VALUES ('%s' , %f ); select ID = convert(bigint, SCOPE_IDENTITY()) ",strings.Title(strings.ToLower(name)),price)


	rows,err:=db.Query(insertProduct_sql)
	if err !=nil{
		f.Println("Error occured while inserting a record", err.Error())
		return -1,err
	}


	defer rows.Close()
	var lastInsertId1 int64
	for rows.Next() {
		rows.Scan(&lastInsertId1)

	}


	return lastInsertId1,err
}

func  InfoMsG(db *sql.DB,id int64)  {
	infoQuery:=f.Sprintf("Select name from Products where id=%d",id)
	rows,err := db.Query(infoQuery)
	if err !=nil{
		f.Println("Error occured while giving info: ", err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var name string
		var id =id
		err:=rows.Scan(&name)
		if err !=nil {
			f.Println("Error reading end process product id with, " , id, err)
		}else{
			f.Printf(name + " product has been created " )
		}



	}


}



func UpdateProduct(db *sql.DB)  {
	f.Print("Please enter product id which you want to change: ")
	var id int
	f.Scanln(&id)

	f.Print("Please enter new product name ")
	var name string
	f.Scanln(&name)

	f.Print("Please enter new product'price ")
	var price float64
	f.Scanln(&price)


	update_query := f.Sprintf("UPDATE Products set name='%s', price=%f where id=%d",name,price,id)

	_, err := db.Exec(update_query)
	if err != nil {
		f.Println("Failed: " + err.Error())

	}


	f.Println("Product informations updated successfully")


}
	


func DeleteProduct(db *sql.DB){
	f.Print("Please enter product id which you want to delete: ")
	var id int
	f.Scanln(&id)

	delete_query:=f.Sprintf("DELETE FROM Products where id=%d",id)
	_, err := db.Exec(delete_query)
	if err != nil {
		f.Println("Failed: " + err.Error())

	}


	f.Println("Product deleted successfully")
}