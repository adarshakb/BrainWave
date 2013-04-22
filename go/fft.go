package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println("Hello world ",math.E)

	var N int = 100
	
	x := make([]float64,N)
	for i := 0; i < len(x); i++ {
		x[i] = 1
	}
	ans := make ([] float64 , N)

	for k := 0; k < N; k++ {
		var c float64 = (-2.0 * math.Pi * float64(k)) / float64(N)
		for i := 0; i < N; i++ {
			ans[k] += x[i] * math.Pow(math.E,c*float64(i))
		}
	}

	fmt.Println(ans)
}