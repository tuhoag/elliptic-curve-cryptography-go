package pedersen

import (
	// "github.com/bwesterb/go-ristretto"
	"testing"
)

func TestGenerateH(t *testing.T) {
	H := GenerateH()

	if H == nil {
		t.Errorf("H (%s) is nil", H)
	}
}
