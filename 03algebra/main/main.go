package main

import (
	_3algebra "algorythms/03algebra"
	"fmt"
	"math"
)

var phi = (1 + math.Sqrt(5)) / 2.0

func main() {
	fmt.Println(_3algebra.PowerLogNFast(2, 9))
	fmt.Println(_3algebra.PowerLogNFast(1.001, 1000))
	fmt.Println(_3algebra.FiboIter(6))
	//var f1, f2 int
	//var val float64
	//fmt.Printf("%g\n", phi)
	//for i := 5; i < 20; i++ {
	//	f1, f2 = _3algebra.FiboIter(i), _3algebra.FiboIter(i-1)
	//	val = _3algebra.PowerLogNFast(phi, i)
	//	val /= math.Sqrt(5)
	//	fmt.Printf("%d / %d ?= %g %d\n", f1, f2, float64(f1)/float64(f2), int(val+0.5))
	//}
}
