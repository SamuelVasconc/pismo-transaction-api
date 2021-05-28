package main

import (
	"log"
	"os"

	"github.com/SamuelVasconc/pismo-transaction-api/cmd"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd := cmd.Server{}
	cmd.Initialization()
	cmd.StartServer()
}
