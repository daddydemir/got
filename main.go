package main

import (
	"log"
	"os"

	"github.com/daddydemir/got/handlers"
	"github.com/joho/godotenv"
)

func main() {

	app := handlers.Urls()

	err := app.ListenTLS(Get("PORT"), Get("CERT_PATH"), Get("KEY_PATH"))
	if err != nil {
		log.Fatal("Server has except: ", err)
	}
}

func Get(key string) string {
	err := godotenv.Load("./prod.env")
	if err != nil {
		log.Fatal("Error loading .env file!", err)
	}
	return os.Getenv(key)
}
