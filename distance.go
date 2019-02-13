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
