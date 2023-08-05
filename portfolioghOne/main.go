package main

import (
	"log"
	"os"
	"portfolioghOne/config"
	"portfolioghOne/server"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("starting mini project")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database...")
	dbHandler := server.InitDatabase(config)
	log.Println(dbHandler)

	log.Println("Initializing HTTP Server!")
	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()

}

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "db_pfghone" + env
	}
	return "db_pfghone"
}
