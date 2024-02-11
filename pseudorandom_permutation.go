package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
)

var (
	PSEUDORANDOM_TEST_KEY []byte = []byte("0123456789abcdef0123456789abcdef")
	PSEUDORANDOM_TEST_IV  []byte = []byte("0123456789abcdef")
)

type PseudorandomPermutation struct {
	key []byte
	iv  []byte
}

func NewPseudorandomPermutation(key []byte, iv []byte) *PseudorandomPermutation {
	prp := PseudorandomPermutation{key: key, iv: iv}
	return &prp
}

func (prp *PseudorandomPermutation) Encrypt(value uint32) []byte {
	plaintext := make([]byte, aes.BlockSize)
	binary.BigEndian.PutUint32(plaintext, value)

	block, err := aes.NewCipher(prp.key)

	if err != nil {
		panic(err)
	}

	encrypter := cipher.NewCBCEncrypter(block, prp.iv)

	encrypted := make([]byte, aes.BlockSize)
	encrypter.CryptBlocks(encrypted, plaintext)

	return encrypted
}

func (prp *PseudorandomPermutation) Decrypt(ciphertext []byte) uint32 {
	block, err := aes.NewCipher(prp.key)

	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, prp.iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = ciphertext[:4]

	value := binary.BigEndian.Uint32(ciphertext)

	return value
}
