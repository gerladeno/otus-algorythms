package _6simplesort

import "algorythms/sortingcommon"

type Shell struct {
	array []int
}

func (sh *Shell) GetSlice() []int {
	return sh.array
}

func (sh *Shell) Run(s string) string {
	var err error
	sh.array, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	sh.sort()
	result := sortingcommon.ReturnString(sh.array)
	return *result
}

func (sh *Shell) sort() {
	var (
		n    = len(sh.array)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if sh.array[j-gap] > sh.array[j] {
					sh.array[j-gap], sh.array[j] = sh.array[j], sh.array[j-gap]
				}
				j = j - gap
			}
		}
	}
}

func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
