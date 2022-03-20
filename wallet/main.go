package main

import (
	"flag"
	"log"

	wlcontrollers "bcserver/wallet/controllers"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	port := flag.Uint("port", 8000, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway")

	app := wlcontrollers.NewWalletServer(uint16(*port), *gateway)
	app.Run()
}