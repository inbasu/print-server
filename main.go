package main

import (
	"net/http"
	"print-server/api/handlers"
	"print-server/database"
)

func main() {
	handler := handlers.PrinterHandler{DB: &database.Database{}}
	handler.DB.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Get)
	mux.HandleFunc("POST /", handler.Post)
	mux.HandleFunc("POST /printlabel", handlers.Print)

	http.ListenAndServe(":8080", mux)
	handler.DB.Close()
}
