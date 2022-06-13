package utils

import (
	ristretto "github.com/bwesterb/go-ristretto"
)

func SplitPoint(targetPoint *ristretto.Point, n int) []*ristretto.Point {
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

func SplitScalar(target *ristretto.Scalar, n int) []*ristretto.Scalar {
	scalars := make([]*ristretto.Scalar, n)

	var sum ristretto.Scalar
	sum.SetZero()
	// sum
	for i := 0; i < n-1; i++ {
		scalars[i] = &ristretto.Scalar{}
		scalars[i].Rand()

		sum.Add(scalars[i], &sum)
		// fmt.Println("sum:", sum, "point:", points[i])
	}
	// fmt.Println("sum:", sum)
	// fmt.Println(points)

	// var temp ristretto.Point
	// temp.Set(targetPoint)
	scalars[n-1] = &ristretto.Scalar{}
	// fmt.Println(points)

	scalars[n-1].Set(target)

	// fmt.Println(points)

	scalars[n-1].Sub(scalars[n-1], &sum)

	return scalars
}
