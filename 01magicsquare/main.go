package main

import (
	"fmt"
	"math"
)

func main() {
	size := 25
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if math.Abs(float64(size-x-y)) < 10 && math.Abs(float64(x-y)) < 10 {
				fmt.Print("*  ")
			} else {
				fmt.Print(".  ")
			}
		}
		fmt.Println("")
	}
}
