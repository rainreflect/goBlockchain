package core

import (
	"testing"
	"time"

	"github.com/rainreflect/gochain/crypto"
	"github.com/rainreflect/gochain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32, prevHash types.Hash) *Block {
	header := &Header{
		Version:   1,
		PrevHash:  prevHash,
		Height:    height,
		Timestamp: time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevHash)
	tx := randomTxWithSign(t)
	b.AddTx(tx)
	assert.Nil(t, b.Sign(privKey))

	return b
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, maxHeight uint32) types.Hash {
	prevHeader, err := bc.GetHeader(maxHeight - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)

}
