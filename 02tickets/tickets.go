package tickets

import (
	"fmt"
	"math"
	"strconv"
)

type Tickets struct {
}

func (t Tickets) Run(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "error"
	}
	count := 0
	limit := int(math.Pow(10, float64(2*n)))
	for i := 0; i < limit; i++ {
		if isLucky(i, 2*n) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func isLucky(n, l int) bool {
	format := fmt.Sprintf("%%0%dd", l)
	nums := []byte(fmt.Sprintf(format, n))
	var sum1, sum2 int
	for _, elem := range nums[0 : l/2] {
		sum1 += int(elem)
	}
	for _, elem := range nums[l/2:] {
		sum2 += int(elem)
	}
	return sum1 == sum2
}
