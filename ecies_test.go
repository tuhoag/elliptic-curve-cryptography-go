package tecc

import (
	"testing"
)

func TestEncryptionAndDecryptionCorrectness(t *testing.T) {
	s1, p1 := GenerateECCKeys()

	message := "hello ECIES"

	cipher, err := ECIESEncrypt(p1, []byte(message))
	if err != nil {
		t.Error(err)
	}

	decryptedData, err := ECIESDecrypt(s1, cipher)
	if err != nil {
		t.Error(err)
	}

	decryptedMessage := string(decryptedData)

	if message != decryptedMessage {
		t.Errorf("Plaintext: %s != decrypted text: %s", message, decryptedMessage)
	}
}
