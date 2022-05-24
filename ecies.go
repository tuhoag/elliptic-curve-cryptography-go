package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	// "encoding/json"
	"fmt"
	"golang.org/x/crypto/hkdf"
	"io"

	"github.com/bwesterb/go-ristretto"
)

func main() {
	s1, p1 := GenerateECCKeys()

	message := "hello world"

	cipher, err := ECIESEncrypt(p1, []byte(message))
	if err != nil {
		fmt.Println(err)
	}

	decryptedData, err := ECIESDecrypt(s1, cipher)
	if err != nil {
		fmt.Println(err)
	}

	decryptedMessage := string(decryptedData)

	if message != decryptedMessage {
		fmt.Printf("Plaintext: %s != decrypted text: %s", message, decryptedMessage)
	} else {
		fmt.Println("Correct")
	}
}

type ECIESCipher struct {
	EncryptedData []byte   `json:"encryptedData"`
	Hash          []byte   `json:"hash"`
	SharedKey     [32]byte `json:"sharedKey"`
}

func GenerateECCKeys() (*ristretto.Scalar, *ristretto.Point) {
	var s ristretto.Scalar
	var p ristretto.Point

	s.Rand()
	p.ScalarMultBase(&s)

	return &s, &p
}

func ECIESEncrypt(publicKey *ristretto.Point, message []byte) (*ECIESCipher, error) {
	// generate a random key
	s, p := GenerateECCKeys()

	// calculate shared key
	var S ristretto.Point
	S.ScalarMult(publicKey, s)

	// derive encryption and hash key
	encryptionKey, hashKey, err := KDF(S.Bytes())

	// encrypt data
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encryptedData := gcm.Seal(nonce, nonce, message, nil)

	// hash data
	h := hmac.New(sha256.New, hashKey)
	h.Write(encryptedData)

	// create cipher object
	cipher := ECIESCipher{
		SharedKey:     p.Bytes(),
		Hash:          h.Sum(nil),
		EncryptedData: encryptedData,
	}

	return &cipher, nil
}

func ECIESDecrypt(privateKey *ristretto.Scalar, eciesCipher *ECIESCipher) ([]byte, error) {
	var S ristretto.Point
	S.SetBytes(&eciesCipher.SharedKey)

	var sharedKey ristretto.Point
	sharedKey.ScalarMult(&S, privateKey)

	encryptionKey, hashKey, err := KDF(sharedKey.Bytes())

	// check hash
	h := hmac.New(sha256.New, hashKey)
	h.Write(eciesCipher.EncryptedData)
	calculatedHash := h.Sum(nil)

	if !hmac.Equal(eciesCipher.Hash, calculatedHash) {
		return nil, fmt.Errorf("Hash is not correct")
	}

	// decrypt data
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(eciesCipher.EncryptedData) < nonceSize {
		return nil, fmt.Errorf("Cipher text length %d < nonce size %s", len(eciesCipher.EncryptedData), nonceSize)
	}

	nonce, cipherData := eciesCipher.EncryptedData[:nonceSize], eciesCipher.EncryptedData[nonceSize:]
	rawData, err := gcm.Open(nil, nonce, cipherData, nil)

	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func KDF(sharedKey []byte) ([]byte, []byte, error) {
	var key1, key2 []byte

	key1 = make([]byte, 32)
	key2 = make([]byte, 32)

	kdf := hkdf.New(sha256.New, sharedKey, nil, nil)

	if _, err := io.ReadFull(kdf, key1); err != nil {
		return nil, nil, err
	}
	if _, err := io.ReadFull(kdf, key2); err != nil {
		return nil, nil, err
	}

	return key1, key2, nil
}
