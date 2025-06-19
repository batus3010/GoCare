package hasher

import (
	"encoding/hex"
	"hash/fnv"
)

type Hasher interface {
	Hash(data string) string
}

// fnvHasher implements Hasher using FNV‑1a 64‑bit.
type fnvHasher struct{}

// NewFNVHasher returns a new FNV‑1a Hasher.
func NewFNVHasher() Hasher {
	return &fnvHasher{}
}

func (h *fnvHasher) Hash(data string) string {
	hasher := fnv.New64a()
	_, _ = hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
