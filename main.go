package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func main() {
	fmt.Println("Go blcokchain server")
}
