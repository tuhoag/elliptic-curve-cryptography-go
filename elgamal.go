package main

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
	// "strings"
	// "bytes"
)

func main() {
	var s1 ristretto.Scalar
	var p1 ristretto.Point

	s1.Rand()
	p1.ScalarMultBase(&s1)

	fmt.Println("s1:", s1)
	fmt.Println("p1:", p1)

	// to encrypt a data
	var secretData ristretto.Point
	secretData.Rand()

	// party 2 send to party 1
	var s2 ristretto.Scalar
	s2.Rand()
	fmt.Println("s2:", s2)

	c1, c2 := Encrypt(&s2, &secretData, &p1)
	fmt.Println("c1:", c1, "c2:", c2)

	// party 1 decrypt
	decryptedData := Decrypt(&s1, c1, c2)

	fmt.Println("secret data:", secretData)
	fmt.Println("decrypted data:", decryptedData)
	// fmt.Println(bytes.Equal(secretData.Bytes(), decryptedData.Bytes()))

	if secretData.Equals(decryptedData) {
		fmt.Println("correct")
	} else {
		fmt.Println("incorrect")
	}
}

func Encrypt(privateKey *ristretto.Scalar, secretData *ristretto.Point, receiverPublicKey *ristretto.Point) (*ristretto.Point, *ristretto.Point) {
	var c1 ristretto.Point
	var c2 ristretto.Point

	c1.ScalarMultBase(privateKey)
	//
	// c1.ScalarMultBase(privateKey)

	c2.ScalarMult(receiverPublicKey, privateKey)
	c2.Add(&c2, secretData)

	return &c1, &c2
}

func Decrypt(privateKey *ristretto.Scalar, c1 *ristretto.Point, c2 *ristretto.Point) *ristretto.Point {
	var sharedKey ristretto.Point
	var rawData ristretto.Point

	sharedKey.ScalarMult(c1, privateKey)

	rawData.Sub(c2, &sharedKey)
	return &rawData
}
