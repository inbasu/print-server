package api

import (
	"encoding/json"
	"fmt"
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
	mask := r.URL.Query().Get("mask")

	data := make(map[string]database.Printer)
	for _, p := range h.DB.Read(mask) {
		data[p.Name] = p
	}
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *PrinterHandler) Post(w http.ResponseWriter, r *http.Request) {
	var p database.Printer
	var jsonData []byte

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&p)
	err := h.DB.Create(p.Name, p.IP)
	if err != nil {
		jsonData, _ = json.Marshal(fmt.Sprintf("{ error: %s}", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}
	jsonData, _ = json.Marshal(p)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *PrinterHandler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *PrinterHandler) Delete(w http.ResponseWriter, r *http.Request) {}
