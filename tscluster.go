package tscluster

type Tscluster struct {
	DistanceFunc DistanceFunc
}

type DistanceFunc func([]float64, []float64) (float64, error)

func NewTscluster(distance DistanceFunc) *Tscluster {
	return &Tscluster{
		DistanceFunc: distance,
	}
}
