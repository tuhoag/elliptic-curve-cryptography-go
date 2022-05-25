package utils

import (
	b64 "encoding/base64"
	"github.com/bwesterb/go-ristretto"
)

// The prime order of the base point is 2^252 + 27742317777372353535851937790883648493.
// var n25519, _ = new(big.Int).SetString("7237005577332262213973186563042994240857116359379907606001950938285454250989", 10)

func ConvertStringToPoint(s string) ristretto.Point {
	bytes, _ := b64.StdEncoding.DecodeString(s)

	point := ConvertBytesToPoint(bytes)
	return point
}

func ConvertStringToScalar(s string) ristretto.Scalar {
	bytes, _ := b64.StdEncoding.DecodeString(s)

	scalar := ConvertBytesToScalar(bytes)

	return scalar
}

func ConvertBytesToPoint(b []byte) ristretto.Point {
	var H ristretto.Point
	var hBytes [32]byte

	copy(hBytes[:32], b[:])

	// result := H.SetBytes(&hBytes)
	// fmt.Println("in convertBytesToPoint result:", result)

	return H
}

func ConvertBytesToScalar(b []byte) ristretto.Scalar {
	var r ristretto.Scalar
	var rBytes [32]byte

	copy(rBytes[:32], b[:])

	// result := r.SetBytes(&rBytes)
	// fmt.Println("in convertBytesToScalar result:", result)

	return r
}

func ConvertScalarToString(scalar ristretto.Scalar) string {
	s := b64.StdEncoding.EncodeToString(scalar.Bytes())
	return s
}

func ConvertPointToString(point ristretto.Point) string {
	s := b64.StdEncoding.EncodeToString(point.Bytes())

	return s
}
