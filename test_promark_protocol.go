package main

import (
	"fmt"
	ristretto "github.com/bwesterb/go-ristretto"
	pedersen "github.com/tuhoag/elliptic-curve-cryptography-go/pedersen"
	utils "github.com/tuhoag/elliptic-curve-cryptography-go/utils"
)

func main() {
	n := 3
	var H ristretto.Point
	H.Rand()

	sVals := make([]ristretto.Scalar, n)
	r1Vals := make([]ristretto.Scalar, n)
	c1Vals := make([]ristretto.Point, n)

	var sumc1 ristretto.Point
	var sumr1 ristretto.Scalar
	sumc1.SetZero()
	sumr1.SetZero()

	for i := 0; i < n; i++ {
		sVals[i].Rand()

		ci, ri := generateCommitment(&H, &sVals[i])

		c1Vals[i] = *ci
		r1Vals[i] = *ri

		sumc1.Add(&sumc1, ci)
		sumr1.Add(&sumr1, ri)
	}

	r2Vals := utils.SplitScalar(&sumr1, n)
	// c2Vals := utils.SplitPoint(sumc1, n)

	var sumc2 ristretto.Point
	sumc2.SetZero()

	for i := 0; i < n; i++ {
		ci := pedersen.CommitTo(&H, r2Vals[i], &sVals[i])
		sumc2.Add(&sumc2, ci)
	}

	fmt.Printf("r1s: %s \n", r1Vals)
	fmt.Printf("r2s: %s \n", r2Vals)

	if sumc1.Equals(&sumc2) {
		fmt.Println("wow")
	} else {
		fmt.Println("fuck")
	}
}

func generateCommitment(H *ristretto.Point, s *ristretto.Scalar) (*ristretto.Point, *ristretto.Scalar) {
	var r ristretto.Scalar
	r.Rand()

	c := pedersen.CommitTo(H, &r, s)
	return c, &r
}
