package tecc

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
	// "strings"
	"testing"
)

func TestSplitPoints(t *testing.T) {
	var targetPoint ristretto.Point
	targetPoint.Rand()
	n := 3

	// fmt.Println("targetPoint:", targetPoint)
	// fmt.Println("n:", n)

	points := Split(&targetPoint, n)

	fmt.Println(points)
	var sum ristretto.Point
	sum.SetZero()

	// fmt.Println("init sum: ", sum)

	for i := 0; i < n; i++ {
		sum.Add(&sum, points[i])
		// fmt.Println("sum:", sum, "point:", points[i])
	}

	if !sum.Equals(&targetPoint) {
		t.Errorf("Sum of generated points: %s is different from the original point: %s", sum, targetPoint)
		// fmt.Println("incorrect")
	}
}
