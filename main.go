package main

import (
	"net/http"
	"print-server/api"
	"print-server/database"
)

func main() {
	handler := api.PrinterHandler{DB: &database.Database{}}
	handler.DB.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.Get)

	http.ListenAndServe(":8080", mux)
	handler.DB.Close()
}
