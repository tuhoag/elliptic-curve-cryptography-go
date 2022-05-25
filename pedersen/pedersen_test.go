package pedersen

import (
	// "github.com/bwesterb/go-ristretto"
	"testing"

	"github.com/bwesterb/go-ristretto"
)

func TestGenerateH(t *testing.T) {
	H := GenerateH()

	if H == nil {
		t.Errorf("H (%s) is nil", H)
	}
}

func TestCommit(t *testing.T) {
	H := GenerateH()

	var s, s1, s2, r, r1, r2 ristretto.Scalar
	s1.Rand()
	r1.Rand()
	s2.Rand()
	r2.Rand()

	C1 := CommitTo(H, &s1, &r1)
	C2 := CommitTo(H, &s2, &r2)

	var sumC ristretto.Point
	sumC.Add(C1, C2)

	s.Add(&s1, &s2)
	r.Add(&r1, &r2)
	C := CommitTo(H, &s, &r)

	if !C.Equals(&sumC) {
		t.Errorf("sum C (%s) != C (%s)", sumC, C)
	}
}
