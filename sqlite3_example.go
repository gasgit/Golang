package main

import (
	"database/sql"	
	"os"
	"fmt"
	 "github.com/mattn/go-sqlite3"
	
)

type Employee struct{
	ID int
	First string
	Last string
	DOB string
	PPS string
}

func main(){
	fmt.Println(" Connected to sqlite3 ")
	db := "staff.db"

	conn, err := sqlite3.Open(db)
	if err != nil{
		fmt.Println("not happening here")
		os.Exit(1)
	}
}