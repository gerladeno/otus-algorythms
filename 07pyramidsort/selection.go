package _7pyramidsort

import "algorythms/sortingcommon"

type Selection struct {
	a []int
}

func (sl *Selection) Run(s string) string {
	var err error
	sl.a, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	sl.sort()
	result := sortingcommon.ReturnString(sl.a)
	return *result
}

func (sl *Selection) GetSlice() []int {
	return sl.a
}

func (sl *Selection) sort() {
	var idx int
	for j := len(sl.a) - 1; j >= 0; j-- {
		idx = getMaxIdx(sl.a[0 : j+1])
		sortingcommon.Swap(sl, idx, j)
	}
}

func getMaxIdx(a []int) int {
	maxIdx := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
