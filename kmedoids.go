package tscluster

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
	"reflect"
)

func (tc *Tscluster) Kmedoids(data [][]float64, k int, maxIterations int) ([]int, error) {
	distMatrix, err := distMatrix(data, tc.DistFunc)
	if err != nil {
		return nil, err
	}
	
	// Pick the initial medoids randomly
	rand.Seed(time.Now().Unix())
	medoids := rand.Perm(len(data))[0:k]
	labels := assignCluster(distMatrix, medoids)

	// iteration for updating medoids
	var updatedLables []int
	for i := 0; i < maxIterations; i++ {
		medoids = findNewMedoids(distMatrix, labels, medoids)
		updatedLables = assignCluster(distMatrix, medoids)
		if reflect.DeepEqual(labels, updatedLables) {
			break
		}
		labels = updatedLables
	}

	return updatedLables, nil
}

func distMatrix(data [][]float64, distFunc DistFunc) (*mat.Dense, error) {
	l := len(data)
	distMatrix := mat.NewDense(l, l, nil)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i == j {
				continue
			} else {
				dist, err := distFunc(data[i], data[j])
				if err != nil {
					return nil, err
				}
				distMatrix.Set(i, j, dist)
			}
		}
	}

	return distMatrix, nil
}

func assignCluster(distMatrix *mat.Dense, medoids []int) []int {
	dataSize, _ := distMatrix.Dims()
	var labels []int
	var medoid int
	for i := 0; i < dataSize; i++ {
		for j, m := range medoids {
			if i == m || j == 0 {
				medoid = m
				continue
			} else if distMatrix.At(i, m) < distMatrix.At(i, medoid) {
				medoid = m
			}
		}
		labels = append(labels, medoid)
	}
	return labels
}

func findNewMedoids(distMatrix *mat.Dense, labels []int, beforeMedoids []int) []int {
	var afterMedoids []int
	for _, m := range beforeMedoids {
		var indexes []int
		var dist []float64
		for i, l := range labels {
			if l == m {
				indexes = append(indexes, i)
			}
		}
		for _, i := range indexes {
			sumDist := 0.
			for _, j := range indexes {
				if i != j {
					sumDist += distMatrix.At(i, j)
				}
			}
			dist = append(dist, sumDist)
		}
		afterMedoids = append(afterMedoids, indexes[minIntSlice(dist)])
	}
	return afterMedoids
}

func minIntSlice(data []float64) int {
	var minIndex int
	var minValue float64
	for i, e := range data {
		if i == 0 || e < minValue {
			minIndex = i
			minValue = e
		}
	}
	return minIndex
}
