package util

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomId() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func RandomColor() string {
	bytes := make([]byte, 3)
	_, err := rand.Read(bytes)
	if err != nil {
		return "#000000"
	}
	return "#" + hex.EncodeToString(bytes)
}
