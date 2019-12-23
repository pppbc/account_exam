package main

import (
	"account_exam/cmd/db"
	"account_exam/cmd/router"
)

func main() {
	db.ConfigDB()
	router.ConfigRouters()
}
