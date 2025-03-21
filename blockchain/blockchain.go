package blockchain

import "time"

var Blockchain []Block

func InitBlockChain() {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		Hash:         "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
}

func AddBlock(transactions []Transaction) {
	prevBlock := Blockchain[len(Blockchain)-1]
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	Blockchain = append(Blockchain, newBlock)
}
