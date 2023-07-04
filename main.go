package main

import (
	"finance_backend/database"
	"finance_backend/routes"
	"fmt"
)

func main() {
	database.ConnectToDataBase()
	routes.HandleRequest()
	fmt.Println("Init server")
}
