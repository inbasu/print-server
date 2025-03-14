package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"print-server/database"
	"regexp"
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

	if !IsValidIp(p.IP) {
		jsonData = []byte("{ \"error\": \"Invalid IP\"}")
		w.WriteHeader(http.StatusBadRequest)
	} else if h.DB.Create(p.Name, p.IP) != nil {
		jsonData = []byte("{ \"error\": \"Database Error\"}")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		jsonData, _ = json.Marshal(p)
		w.WriteHeader(http.StatusCreated)
	}
	w.Write(jsonData)
}

func (h *PrinterHandler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *PrinterHandler) Delete(w http.ResponseWriter, r *http.Request) {}

/* Скорее всего это утилья */
func IsValidIp(ip string) bool {
	ipPattern, _ := regexp.Compile("^((?:[01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[01]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])$")
	return ipPattern.Match([]byte(ip))
}
