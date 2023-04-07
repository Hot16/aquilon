package main

import (
	"aquilon/internal"
	send_request "aquilon/pkg"
	"aquilon/pkg/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	response := send_request.SendRequest()
	clients := internal.TransformClients(response)
	fmt.Println(clients)

	db, err := database.NewDatabase()
	fmt.Println("DB ", db)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	batchSize, err := strconv.Atoi(os.Getenv("BATCH_SIZE"))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	err = db.SaveClients(clients, batchSize)
	if err != nil {
		log.Fatalf("Failed to save clients to database: %v", err)
	}
}
