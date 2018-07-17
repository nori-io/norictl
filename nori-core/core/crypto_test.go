package core

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFromString(t *testing.T) {
	assert := assert.New(t)

	pk, err := NewPasskey()
	assert.NotNil(pk)
	assert.Nil(err)

	pkString := pk.String()

	fromStr, err := PasskeyFromString(pkString)
	assert.NotNil(pk)
	assert.Nil(err)

	assert.Equal(pk.bs, fromStr.bs)
}

func TestCrypto(t *testing.T) {
	assert := assert.New(t)

	pk, err := NewPasskey()
	assert.NotNil(pk)
	assert.Nil(err)

	testData := make([]byte, 512)
	_, err = rand.Read(pk.bs)
	assert.Nil(err)

	ciphertext, hash, err := pk.Encrypt(testData)
	assert.NotEmpty(ciphertext)
	assert.NotEmpty(hash)
	assert.Nil(err)

	data, err := pk.Decrypt(ciphertext, hash)
	assert.NotEmpty(data)
	assert.Nil(err)

	assert.Equal(testData, data)

	ciphertext[256] = uint8(ciphertext[256] * 2)

	data, err = pk.Decrypt(ciphertext, hash)
	assert.Nil(data)
	assert.Equal(err.Error(), "malformed signature")
}
