package routers

import (
	"net/http"
	"print-server/api/handlers"
)

type Handler struct {
	*handlers.PrinterHandler
}

func (h *Handler) RegisterPrinterRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /run", h.Get)
	mux.HandleFunc("POST /run", h.Post)
	mux.HandleFunc("POST /printlabel", handlers.Print)
	return mux
}
