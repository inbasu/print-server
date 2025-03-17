package main

import (
	"net/http"
	routers "print-server/api"
	"print-server/api/handlers"
	"print-server/database"
)

func main() {
	handler := handlers.PrinterHandler{DB: &database.Database{}}
	handler.DB.Connect()

	printerHandler := routers.Handler{&handler}
	printerRouter := printerHandler.RegisterPrinterRouter()

	mux := http.NewServeMux()
	mux.Handle("/api/printer/", http.StripPrefix("/api/printer", printerRouter))

	http.ListenAndServe(":8181", mux)
	handler.DB.Close()
}
