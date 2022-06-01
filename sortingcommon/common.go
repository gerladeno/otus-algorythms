package sortingcommon

import (
	"errors"
	"strconv"
	"strings"
)

func PrepareSlice(s *string) ([]int, error) {
	ss := strings.Split(*s, "\r\n")
	if len(ss) < 2 {
		return nil, errors.New("broken file")
	}
	l, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil, err
	}
	a := make([]int, 0, l)
	var tmp int
	for _, elem := range strings.Fields(ss[1]) {
		tmp, err = strconv.Atoi(elem)
		if err != nil {
			return nil, err
		}
		a = append(a, tmp)
	}
	return a, err
}

func ReturnString(a []int) *string {
	ss := make([]string, 0, len(a))
	for _, elem := range a {
		ss = append(ss, strconv.Itoa(elem))
	}
	result := strings.Join(ss, " ")
	return &result
}

type Sorter interface {
	GetSlice() []int
}

func Swap(s Sorter, i, j int) {
	a := s.GetSlice()
	a[i], a[j] = a[j], a[i]
}

func BinarySearch(s Sorter, left, right, n int) int {
	a := s.GetSlice()
	return binarySearch(a, left, right, n)
}

func binarySearch(a []int, left, right, n int) int {
	m := (left + right) / 2
	if a[m] == n || left == right {
		return m
	}
	if left+1 == right {
		if a[left] < n {
			return right
		}
		return left
	}
	if a[m] < n {
		return binarySearch(a, m, right, n)
	}
	return binarySearch(a, left, m, n)
}
