package core

import (
	"fmt"

	"github.com/rainreflect/gochain/crypto"
	"github.com/rainreflect/gochain/types"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature

	//cached tx data hash
	hash types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Hash(h Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = h.Hash(tx)
	}
	return h.Hash(tx)
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sign, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sign

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
