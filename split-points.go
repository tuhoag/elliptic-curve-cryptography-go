package main

import (
	"fmt"
	"github.com/bwesterb/go-ristretto"
	// "strings"
)

func main() {
	var targetPoint ristretto.Point
	targetPoint.Rand()
	n := 3

	fmt.Println("targetPoint:", targetPoint)
	fmt.Println("n:", n)

	points := Split(&targetPoint, n)

	fmt.Println(points)
	var sum ristretto.Point
	sum.SetZero()

	fmt.Println("init sum: ", sum)

	for i := 0; i < n; i++ {
		sum.Add(&sum, points[i])
		fmt.Println("sum:", sum, "point:", points[i])
	}

	if sum.Equals(&targetPoint) {
		fmt.Println("correct")
	} else {
		fmt.Println("incorrect")
	}
}

func Split(targetPoint *ristretto.Point, n int) []*ristretto.Point {
	points := make([]*ristretto.Point, n)

	var sum ristretto.Point
	sum.SetZero()
	// sum
	for i := 0; i < n-1; i++ {
		points[i] = &ristretto.Point{}
		points[i].Rand()

		sum.Add(points[i], &sum)
		// fmt.Println("sum:", sum, "point:", points[i])
	}
	// fmt.Println("sum:", sum)
	// fmt.Println(points)

	// var temp ristretto.Point
	// temp.Set(targetPoint)
	points[n-1] = &ristretto.Point{}
	// fmt.Println(points)

	points[n-1].Set(targetPoint)

	// fmt.Println(points)

	points[n-1].Sub(points[n-1], &sum)

	return points
}
