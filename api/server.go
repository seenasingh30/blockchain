package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/seenasingh30/blockchain/blockchain"

	"github.com/gorilla/mux"
)

// Get Blockchain
func getBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain.Blockchain)
}

// Create Transaction
func createTransaction(w http.ResponseWriter, r *http.Request) {
	var tx blockchain.Transaction
	_ = json.NewDecoder(r.Body).Decode(&tx)

	// Add to blockchain
	blockchain.AddBlock([]blockchain.Transaction{tx})

	json.NewEncoder(w).Encode(tx)
}

// Start HTTP Server
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/blocks", getBlockchain).Methods("GET")
	router.HandleFunc("/transaction", createTransaction).Methods("POST")

	fmt.Println("API Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
