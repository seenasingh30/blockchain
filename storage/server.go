package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// Save blockchain to file
func SaveBlockchain(data interface{}, filename string) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(filename, file, 0644)
	fmt.Println("Blockchain saved!")
}
