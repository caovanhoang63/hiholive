package core

import (
	"crypto/sha256"
	"encoding/hex"
)

type sha256Hash struct{}

func NewSha256Hash() *sha256Hash {
	return &sha256Hash{}
}

func (h *sha256Hash) Hash(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
