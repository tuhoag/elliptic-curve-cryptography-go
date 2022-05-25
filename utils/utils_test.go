package utils

import (
	ristretto "github.com/bwesterb/go-ristretto"
	"testing"
)

func TestConvertStringToPoint(t *testing.T) {
	// create a point
	var p ristretto.Point

	// convert it to string
	pstr := ConvertPointToString(&p)

	// convert to string to point
	convertedP, err := ConvertStringToPoint(pstr)

	if err != nil {
		t.Error(err)
	}

	// they must be equal
	if !convertedP.Equals(&p) {
		t.Errorf("Converted point (%s) != raw point (%s)\n", convertedP, p)
	}
}
