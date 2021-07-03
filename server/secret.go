package server

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"pravaah/config"
)

func InitSecret() error {
	randBytes := make([]byte, 32)

	// Generate a random string for secret
	if _, err := rand.Read(randBytes); err != nil {
		fmt.Printf("Unable to generate secret. Exiting.\n")
		return err
	}

	// Encode a hex string
	config.Secret = hex.EncodeToString(randBytes)

	return nil
}
