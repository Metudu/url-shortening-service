package main

import (
	"github.com/Metudu/url-shortening-service/db"
	"github.com/Metudu/url-shortening-service/server"

	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()	
	defer db.GetDB().Close()
	server.StartServer(server.InitializeServer())
}