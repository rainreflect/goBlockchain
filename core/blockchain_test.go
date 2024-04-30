package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockChainwWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockChain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockChainwWithGenesis(t)

	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestAddBlock(t *testing.T) {
	bc := newBlockChainwWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))

		assert.Nil(t, bc.AddBlock(block))
	}
	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	assert.NotNil(t, bc.AddBlock((randomBlock(89))))
}
func TestHasBlock(t *testing.T) {
	bc := newBlockChainwWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}
