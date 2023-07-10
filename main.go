package main

import (
	"fmt"
	database "user_auth/database/config"
	routes "user_auth/routes"
)

type Main struct {
	route routes.Route
}

func main() {
	database.ConnectToDataBase()
	fmt.Println("Init server")
	handle := Main{route: routes.Route{}}
	handle.route.HandleRequest()
}
