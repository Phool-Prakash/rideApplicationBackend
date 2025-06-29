package server

import (
	"log"
	"net/http"
	"rideApplicationBackend/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.ConnectDB()
	defer db.Disconnect(nil)

	// router :=
	log.Println("Server Started at :8080")
	log.Fatal(http.ListenAndServe(":8080",router))
}