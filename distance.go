package tscluster

import (
	"math"
	"github.com/golang/go/src/errors"
)

func EuclideanDistance(firstVec, secondVec []float64) (float64, error) {
	distance := 0.
	if len(firstVec) != len(secondVec) {
		return 0., errors.New("When measuring the Euclidean distance, the size of the vector must be the same.")
	}
	for i := range firstVec {
		distance += (firstVec[i] - secondVec[i]) * (firstVec[i] - secondVec[i])
	}
	return math.Sqrt(distance), nil
}

func euclideanDistance(x, y float64) float64 {
	difference := x - y
	return math.Sqrt(difference * difference)
}

func DTW(firstVec, secondVec []float64) (float64, error) {
	n := len(firstVec)
	m := len(secondVec)
	mat := make([]float64, n*m)

	mat[0] = euclideanDistance(firstVec[0], secondVec[0])
	for x := 1; x < n; x++ {
		mat[x] = mat[x-1] + euclideanDistance(firstVec[x], secondVec[0])
	}
	off := n
	for y := 1; y < m; y++ {
		mat[off] = mat[off-n] + euclideanDistance(firstVec[0], secondVec[y])
		off++
		for x := 1; x < n; x++ {
			minCost := min(mat[off-1], min(mat[off-n], mat[off-n-1]))
			mat[off] = minCost + euclideanDistance(firstVec[x], secondVec[y])
			off++
		}
	}

	path := make([]Point, 0)
	x, y := n-1, m-1
	path = append(path, Point{X: x, Y: y})
	for x > 0 || y > 0 {
		o := y*n + x
		diag := math.Inf(1)
		left := math.Inf(1)
		down := math.Inf(1)
		if x > 0 && y > 0 {
			diag = mat[o-n-1]
		}
		if x > 0 {
			left = mat[o-1]
		}
		if y > 0 {
			down = mat[o-n]
		}
		switch {
		case diag <= left && diag <= down:
			x--
			y--
		case left < diag && left < down:
			x--
		case down < diag && down < left:
			y--
		case x <= y:
			x--
		default:
			y--
		}
		path = append(path, Point{X: x, Y: y})
	}
	
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - i - 1
		path[i], path[j] = path[j], path[i]
	}

	return mat[n*m-1], nil
}
