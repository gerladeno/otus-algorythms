package _8quicksort

import (
	"algorythms/sortingcommon"
)

type Merge struct {
	Ch chan struct{}
	a  []int
}

func (m *Merge) Run(s string) string {
	var err error
	m.a, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	m.sort()
	result := sortingcommon.ReturnString(m.a)
	return *result
}

func (m *Merge) GetSlice() []int {
	return m.a
}

func (m *Merge) sort() {
	m.mergeSort(0, len(m.a)-1)
}

func (m *Merge) mergeSort(l, r int) {
	if l >= r {
		return
	}
	M := (l + r) / 2
	m.mergeSort(l, M)
	m.mergeSort(M+1, r)
	m.merge(l, M, r)
}

func (m *Merge) merge(l, x, r int) {
	M := make([]int, 0, r-l+1)
	L := l
	R := x + 1
	for L <= x && R <= r {
		if m.a[L] < m.a[R] {
			M = append(M, m.a[L])
			L++
		} else {
			M = append(M, m.a[R])
			R++
		}
	}
	for L <= x {
		M = append(M, m.a[L])
		L++
	}
	for R <= r {
		M = append(M, m.a[R])
		R++
	}
	for i := l; i <= r; i++ {
		m.a[i] = M[i-l]
	}
}
