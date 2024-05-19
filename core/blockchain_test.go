package core

import (
	"fmt"
	"testing"

	"github.com/rainreflect/gochain/types"
	"github.com/stretchr/testify/assert"
)

func newBlockChainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockChain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestAddBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		maxHeight := uint32(i + 1)
		block := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, maxHeight))

		assert.Nil(t, bc.AddBlock(block))
	}
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	assert.NotNil(t, bc.AddBlock((randomBlock(89, types.Hash{}))))
}
func TestHasBlock(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
	assert.False(t, bc.HasBlock(100))
}

func TestAddBlockToHigh(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 0, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockChainWithGenesis(t)

	lenBlocks := 2
	for i := 0; i < lenBlocks; i++ {
		maxHeight := uint32(i + 1)
		block := randomBlockWithSignature(t, maxHeight, getPrevBlockHash(t, bc, maxHeight))
		assert.Nil(t, bc.AddBlock(block))

	}
}
