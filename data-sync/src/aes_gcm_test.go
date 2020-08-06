package main

import (
	"encoding/hex"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAesGCMEncrypt(t *testing.T) {
	Convey("Testing of aes-gcm encrypt", t, func() {
		keyHex := []byte("af5c5d660ca7a22cdbedf09465c191492a4ea99695bec34f136d8a15bd3616d2")
		nonceHex := []byte("fa8b0eb0aac5ad2331449bbc")
		plaintext := []byte("just a test")
		wantsHex := "3c1edb8875c4defa9787d9e7b4d0b23bfb30a1e8b60e48049a64ce"
		Convey("encrypt should pass", func() {
			key, _ := hex.DecodeString(string(keyHex))
			nonce, _ := hex.DecodeString(string(nonceHex))
			ret, err := AesGCMEncrypt(key, nonce, plaintext)
			So(err, ShouldEqual, nil)
			Convey("encrypt result should equal", func() {
				So(hex.EncodeToString(ret), ShouldEqual, wantsHex)
			})
		})
	})
}

func TestAesGCMDecrypt(t *testing.T) {
	Convey("Testing of aes-gcm decrypt", t, func() {
		keyHex := []byte("af5c5d660ca7a22cdbedf09465c191492a4ea99695bec34f136d8a15bd3616d2")
		nonceHex := []byte("fa8b0eb0aac5ad2331449bbc")
		cipherHex := []byte("3c1edb8875c4defa9787d9e7b4d0b23bfb30a1e8b60e48049a64ce")
		wants := "just a test"
		Convey("Decrypt should pass", func() {
			key, _ := hex.DecodeString(string(keyHex))
			nonce, _ := hex.DecodeString(string(nonceHex))
			cipherText, _ := hex.DecodeString(string(cipherHex))
			ret, err := AesGCMDecrypt(key, nonce, cipherText)
			So(err, ShouldEqual, nil)
			Convey("decrypt result should equal", func() {
				So(string(ret), ShouldEqual, wants)
			})
		})
	})
}
