package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Pass(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	//adr := publicKey.Address()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)

	assert.True(t, sig.Verify(publicKey, msg))

}

func TestKeypair_Sign_Verify_Fail(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	//adr := publicKey.Address()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	newPrivKey := GeneratePrivateKey()
	newPublicKey := newPrivKey.PublicKey()

	assert.False(t, sig.Verify(newPublicKey, msg))
	assert.False(t, sig.Verify(publicKey, []byte("hero workld")))
}
