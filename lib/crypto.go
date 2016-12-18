package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"golang.org/x/crypto/scrypt"
	"log"
)

func Decrypt(key, nonce, ciphertext []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, err
}

func Encrypt(key, nonce, cleartext []byte) (ciphertext []byte, err error) {
	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, cleartext, nil)

	return ciphertext, err
}

func NewNonce() (nonce []byte) {
	nonce = make([]byte, 12)
	rand.Read(nonce)
	return nonce
}

func NewKey(salt, password []byte) (key []byte) {
	key, err = scrypt.Key(password, salt, 16384, 8, 1, 32)

	if err != nil {
		log.Fatal(err)
	}

	return key
}
