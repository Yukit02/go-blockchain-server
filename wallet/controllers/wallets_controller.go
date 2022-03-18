package wlcontrollers

import (
	"log"
	"net/http"
	"strconv"
)

type WalletServer struct {
	Port uint16
	Gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{Port: port, Gateway: gateway}
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Printf("ERROR: Invalid HTTP Method")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+ strconv.Itoa(int(ws.Port)), nil))
}
