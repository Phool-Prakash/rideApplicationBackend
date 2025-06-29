package routes

import (
	"rideApplicationBackend/internal/handlers"

	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router{
	r:= mux.NewRouter()
	r.HandleFunc("send-otp",handlers.Send)
}