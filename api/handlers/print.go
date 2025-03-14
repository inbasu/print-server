package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type printData struct {
	ip  string
	zpl string
}

func Print(w http.ResponseWriter, r *http.Request) {
	var d printData
	json.NewDecoder(r.Body).Decode(&d)
	conn, _ := net.Dial("tcp", d.ip+"9100")
	fmt.Fprintf(conn, d.zpl)
}
