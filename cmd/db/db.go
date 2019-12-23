package db

import (
	_ "github.com/bmizerany/pq"
	"github.com/jmoiron/sqlx"
	"log"
)

//全局的DB对象
var DB *sqlx.DB

//初始化数据库
func ConfigDB() {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres password=123456 dbname=postgres")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Success Init PostgresSQL")
}
