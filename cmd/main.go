package main

import (
	"log"

	"github.com/Zero-cold050903/go-project/cmd/api"
	"github.com/Zero-cold050903/go-project/config"
	"github.com/Zero-cold050903/go-project/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Password:             config.Envs.DBPassword,
		Host:                 config.Envs.DBHost,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
