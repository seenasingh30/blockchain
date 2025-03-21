package main

import (
	"fmt"

	"github.com/seenasingh30/blockchain/api"
	"github.com/seenasingh30/blockchain/blockchain"
)

func main() {
	blockchain.InitBlockChain()
	api.StartServer()

	fmt.Println("Blockchain Server running...")

}
