package main

import (
	database "account_exam/cmd/db"
	"account_exam/cmd/router"
)

func main() {
	//database.ConnectDB()
	router.ConfigRouters()
	defer database.DisconnectedDB()
}
