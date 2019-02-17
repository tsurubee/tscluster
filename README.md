# tscluster
tscluster is time-series clustering implemented in Go.  

## Usage
### Installation
```
$ go get github.com/tsurubee/tscluster
```

### Example
```go
func main() {
	var dataset [][]float64
	dataset = append(dataset, []float64{1, 1, 1, 1, 1, 1})
	dataset = append(dataset, []float64{1, 1, 1, 1, 1, 1, 1})
	dataset = append(dataset, []float64{2, 2, 2, 2, 2})
	dataset = append(dataset, []float64{2, 2, 2, 2, 2, 2, 2})

	tc := tscluster.NewTscluster(tscluster.DTW)
	labels, err := tc.Kmedoids(dataset, 2, 20)
	if err != nil {
	log.Fatal(err)
	}
	fmt.Println(labels)
}

#=>
[0 0 2 2]
```
It returns the label of the result of clustering on the given data array.  
The numerical value of the label returns the index number of medoid of the cluster to which the data belongs.  

## License

[MIT](https://github.com/tsurubee/tscluster/blob/master/LICENSE)

## Author

[tsurubee](https://github.com/tsurubee)