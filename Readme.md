# Blockchain Project in Golang

## Overview
This project is a simple blockchain implementation in Golang, featuring transactions, wallets, proof-of-work, and a peer-to-peer network.

## Features
- Blockchain with proof-of-work (PoW)
- Transactions with digital signatures
- Wallet system (public/private key pairs)
- P2P networking
- REST API for interacting with the blockchain

## Project Structure
```
blockchain-project/
│── main.go
│── go.mod
│── go.sum
│
├── blockchain/
│   ├── blockchain.go
│   ├── block.go
│   ├── proof_of_work.go
│   ├── transaction.go
│
├── wallet/
│   ├── wallet.go
│
├── network/
│   ├── p2p.go
│
├── api/
│   ├── server.go
```

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/blockchain-project.git
   cd blockchain-project
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the project:
   ```sh
   go run main.go
   ```

## API Endpoints
- **Get Blockchain:** `GET /blocks`
- **Create Transaction:** `POST /transaction`
  ```json
  {
    "Sender": "Alice",
    "Receiver": "Bob",
    "Amount": 10
  }
  ```


