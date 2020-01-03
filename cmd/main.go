package main

import (
	_ "account_exam/cmd/app"
	database "account_exam/cmd/db"
	"account_exam/cmd/router"
)

func main() {
	router.ConfigRouters()
	defer database.DisconnectedDB()
}
