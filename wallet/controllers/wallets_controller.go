package wlcontrollers

import (
	"io"
	"log"
	"net/http"
	"strconv"

	wlmodels "bcserver/wallet/models"
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

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "spplication/json")
		myWallet, err := wlmodels.NewWallet()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		m, err := myWallet.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/", ws.Wallet)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+ strconv.Itoa(int(ws.Port)), nil))
}
