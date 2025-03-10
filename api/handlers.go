package api

import (
	"encoding/json"
	"net/http"
	"print-server/database"
)

type Handler interface {
	Get()
}

type PrinterHandler struct {
	DB *database.Database
}

func (h *PrinterHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mask := r.URL.Query().Get("mask")

	data := make(map[string]database.Printer)
	for _, p := range h.DB.Read(mask) {
		data[p.Name] = p
	}
	jsonData, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *PrinterHandler) Post(w http.ResponseWriter, r *http.Request) {}

func (h *PrinterHandler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *PrinterHandler) Delete(w http.ResponseWriter, r *http.Request) {}
