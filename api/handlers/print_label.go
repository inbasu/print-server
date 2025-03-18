package handlers

import (
	"encoding/json"
	"net"
	"net/http"
)

type printData struct {
	IP  string `json:"ip"`
	ZPL string `json:"zpl"`
}

func writeResponse(w http.ResponseWriter, status int, data string) {
	w.WriteHeader(status)
	w.Write([]byte(data))
}

func Print(w http.ResponseWriter, r *http.Request) {
	var d printData

	w.Header().Set("Content-Type", "application/json")
	if json.NewDecoder(r.Body).Decode(&d) != nil {
		writeResponse(w, http.StatusBadRequest, "{\"error\": \"Bad data\"}")
		return
	}
	if !IsValidIp(d.IP) || false {
		writeResponse(w, http.StatusBadRequest, "{\"error\":  \"Invalid IP\"}")
		return
	}
	dialer := net.Dialer{Timeout: 5000000}
	conn, conErr := dialer.Dial("tcp", d.IP+":9100")
	if conErr != nil {
		writeResponse(w, http.StatusOK, "{\"error\": \"Perhaps printer is offline.\"}")
		return
	}
	conn.Write([]byte(d.ZPL))
	defer conn.Close()
	writeResponse(w, http.StatusOK, "{\"ok\": \"ok\", \"error\": \"\"}")
}
