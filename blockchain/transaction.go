package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

// Transaction structure
type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Signature string
}

// Sign the transaction
func (t *Transaction) SignTransaction(privateKey *ecdsa.PrivateKey) {
	data := fmt.Sprintf("%s%s%f", t.Sender, t.Receiver, t.Amount)
	hash := sha256.Sum256([]byte(data))

	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	signature := append(r.Bytes(), s.Bytes()...)
	t.Signature = hex.EncodeToString(signature)
}

// Verify transaction signature
func (t *Transaction) VerifyTransaction(publicKey ecdsa.PublicKey) bool {
	data := fmt.Sprintf("%s%s%f", t.Sender, t.Receiver, t.Amount)
	hash := sha256.Sum256([]byte(data))

	signatureBytes, _ := hex.DecodeString(t.Signature)
	r := big.Int{}
	s := big.Int{}
	r.SetBytes(signatureBytes[:len(signatureBytes)/2])
	s.SetBytes(signatureBytes[len(signatureBytes)/2:])

	return ecdsa.Verify(&publicKey, hash[:], &r, &s)
}
