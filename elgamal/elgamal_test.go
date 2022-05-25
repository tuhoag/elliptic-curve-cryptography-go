package elgamal

import (
	"github.com/bwesterb/go-ristretto"
	"testing"
)

func TestEncryptionAndDecryption(t *testing.T) {
	var s1 ristretto.Scalar
	var p1 ristretto.Point

	s1.Rand()
	p1.ScalarMultBase(&s1)

	// to encrypt a data
	var secretData ristretto.Point
	secretData.Rand()

	// party 2 send to party 1
	var s2 ristretto.Scalar
	s2.Rand()

	c1, c2 := Encrypt(&s2, &secretData, &p1)

	// party 1 decrypt
	decryptedData := Decrypt(&s1, c1, c2)

	if !secretData.Equals(decryptedData) {
		t.Errorf("Decrypted point (%s) is different from the original point (%s)", decryptedData, secretData)
	}
}
