package main

import (
	"github.com/danielpk74/tweettor/db"
	"github.com/danielpk74/tweettor/handlers"
	"log"
)

func main() {
	conn := db.Conn
	if !conn.CheckConnection() {
		log.Fatalln("MongoDB Server is not available.")
		return
	}
	handlers.Handlers()

}
