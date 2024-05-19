package core

import (
	"crypto/sha256"

	"github.com/rainreflect/gochain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(h *Header) types.Hash {

	hash := sha256.Sum256(h.Bytes())
	return types.Hash(hash)
}
