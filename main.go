package main

import (
	"avito_test_task/db"
	"fmt"
)

func main() {
	db.InitDatabase()
	fmt.Println("Test success")
}
