package _3algebra

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Power struct {
}

func (p Power) Run(s string) string {
	items := strings.Fields(s)
	if len(items) != 2 {
		return "error"
	}
	val, err := strconv.ParseFloat(items[0], 64)
	if err != nil {
		return "error"
	}
	power, err := strconv.Atoi(items[1])
	if err != nil {
		return "error"
	}
	res := math.Round(PowerN(val, power)*100_000_000_000) / 100_000_000_000
	//res := PowerLogN(val, power)
	result := fmt.Sprintf("%g", res)
	if float64(int(res)) == res {
		result = fmt.Sprintf("%.1f", res)
	}
	return result
}

func PowerN(val float64, power int) float64 {
	result := 1.0
	for i := 1; i <= power; i++ {
		result *= val
	}
	return result
}

func PowerLogN(val float64, power int) float64 {
	tmp := fmt.Sprintf("%b", power)
	p := make([]int, 0, len(tmp))
	var v int
	for _, symbol := range tmp {
		v, _ = strconv.Atoi(string(symbol))
		p = append(p, v)
	}
	result := 1.0
	for i, elem := range p {
		if elem == 1 {
			result *= power2(val, len(p)-i)
		}
	}
	return result
}

func power2(val float64, power int) float64 {
	for i := 1; i < power; i++ {
		val *= val
	}
	return val
}

func PowerLogNFast(val float64, power int) float64 {
	result := 1.0
	for power >= 1 {
		if power%2 == 1 {
			result *= val
		}
		power /= 2
		val *= val
	}
	return result
}
