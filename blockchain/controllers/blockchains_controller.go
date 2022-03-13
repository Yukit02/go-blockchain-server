package bccontrollers

import (
	bcmodels "bcserver/blockchain/models"
	wlmodels "bcserver/wallet/models"

	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*bcmodels.Blockchain = make(map[string] *bcmodels.Blockchain)

type BlockchainServer struct {
	Port uint16
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{Port: port}
}

func (bcs *BlockchainServer) GetBlockchain() *bcmodels.Blockchain {
	bc, ok := cache["blockchain"]

	if !ok {
		minerWallet, _ := wlmodels.NewWallet()
		bc = bcmodels.NewBlockchain(minerWallet.BlockchainAddress, bcs.Port)
		cache["blockchain"] = bc
	}

	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+ strconv.Itoa(int(bcs.Port)), nil))
}