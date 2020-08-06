package main


import (
	"crypto/aes"
	"crypto/cipher"
)

//AesGCMEncrypt aes-gcm mode encrypt
func AesGCMEncrypt(key, nonce, plaintext []byte) (ciphertext []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)

	return
}

//AesGCMDecrypt aes-gcm mode decrypt
func AesGCMDecrypt(key, nonce, ciphertext []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err = aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return
}
