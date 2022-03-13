package main

import (
	"flag"
	"log"

	bccontrollers "bcserver/blockchain/controllers"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()

	log.Println(port)
	app := bccontrollers.NewBlockchainServer(uint16(*port))
	app.Run()
}