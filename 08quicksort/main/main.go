package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(mergeSorted(a, b))
}

func mergeSorted(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	result := make([]int, 0, len(a)+len(b))
	if a[0] >= b[len(b)-1] {
		result = append(b, a...)
		return result
	}
	if b[0] >= a[len(a)-1] {
		result = append(a, b...)
		return result
	}
	var i, j int
	for {
		if a[i] >= b[j] {
			result = append(result, b[j])
			j++
			if j == len(b) {
				result = append(result, a[i:]...)
				return result
			}
		} else {
			result = append(result, a[i])
			i++
			if i == len(a) {
				result = append(result, b[j:]...)
				return result
			}
		}
	}
}
