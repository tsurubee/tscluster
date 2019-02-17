package tscluster

type Point struct {
	X, Y int
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

