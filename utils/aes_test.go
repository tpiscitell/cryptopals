package utils

import (
	"testing"
)

func TestAesECB(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	plaintext := []byte("The beatles wrote this song!!!!!")

	out := AesEcbDecrypt(AesEcbEncrypt(plaintext, key), key)
	if string(out) != string(plaintext) {
		t.Error("Encrypted and decrypting with the same key got different results")
	}
}

func TestAesCBC(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	plaintext := []byte("The beatles wrote this song!!!!!")
	iv := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	out := AesCbcDecrypt(AesCbcEncrypt(plaintext, key, iv), key, iv)
	if string(out) != string(plaintext) {
		t.Error("Encrypted and decrypting with the same key got different results")
	}
}

func TestRepeats(t *testing.T) {
	in := []byte{0x00, 0x01, 0x00, 0x01, 0x02, 0x33, 0x45, 0x44, 0x45, 0x44}

	count := CountRepeats(in, 2)

	if count != 4 {
		t.Error("Count should be 4. Got:", count)
	}
}
