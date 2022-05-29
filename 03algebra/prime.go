package _3algebra

import (
	"fmt"
	"strconv"
)

type Prime struct {
}

func (p Prime) Run(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "error"
	}
	result := primeNumber(n)
	return fmt.Sprintf("%d", result)
}

func primeNumber(n int) int {
	var dividers, primes []int
	dividers = make([]int, n+2)
	for i := 2; i <= n; i++ {
		if dividers[i] == 0 {
			dividers[i] = i
			primes = append(primes, i)
		}
		for j := 0; j < len(primes) && primes[j] <= dividers[i] && primes[j]*i <= n; j++ {
			dividers[primes[j]*i] = primes[j]
		}
	}
	return len(primes)
}
