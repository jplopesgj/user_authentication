package main

import (
	database "finance_backend/database/config"
	"fmt"
)

func main() {
	database.ConnectToDataBase()
	fmt.Println("Init server")
}
