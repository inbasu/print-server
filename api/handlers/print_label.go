package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type printData struct {
	IP  string `json:"ip"`
	ZPL string `json:"zpl"`
}

func Print(w http.ResponseWriter, r *http.Request) {
	var d printData
	var jsonData []byte

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&d)
	if !IsValidIp(d.IP) || false {
		jsonData = []byte("{ \"error\": \"Invalid IP\"}")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonData)
		return
	}
	conn, _ := net.Dial("tcp", d.IP+"9100")
	fmt.Fprintf(conn, d.ZPL)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	defer conn.Close()
}
