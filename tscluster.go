package tscluster

type Tscluster struct {
	DistFunc DistFunc
}

type DistFunc func([]float64, []float64) (float64, error)

func NewTscluster(distance DistFunc) *Tscluster {
	return &Tscluster{
		DistFunc: distance,
	}
}
