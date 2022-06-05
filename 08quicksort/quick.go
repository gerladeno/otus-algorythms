package _8quicksort

import (
	"algorythms/sortingcommon"
	"algorythms/visual"
)

type Quick struct {
	Ch chan struct{}
	a  []int
}

func (q *Quick) Run(s string) string {
	var err error
	q.a, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	q.sort()
	result := sortingcommon.ReturnString(q.a)
	return *result
}

func (q *Quick) GetSlice() []int {
	return q.a
}

func (q *Quick) sort() {
	q.quickSort(0, len(q.a)-1)
}

func (q *Quick) quickSort(l, r int) {
	if l >= r {
		return
	}
	x := q.segregate(l, r)
	q.quickSort(l, x-1)
	q.quickSort(x+1, r)
}

func (q *Quick) split(l, r int) int {
	m := l - 1
	for i := l; i <= r; i++ {
		if q.a[i] <= q.a[r] {
			m++
			<-q.Ch
			visual.SwapEvent(q, visual.Swap{I: m, J: i})
			//sortingcommon.Swap(q, m, i)
		}
	}
	return m
}

func (q *Quick) segregate(l, r int) int {
	i := l
	j := r
	val := q.a[r]
	for {
		if i >= j {
			<-q.Ch
			visual.SwapEvent(q, visual.Swap{I: i, J: r})
			return i
		}
		if q.a[i] <= val {
			i++
			continue
		}
		if q.a[j] >= val {
			j--
			continue
		}
		<-q.Ch
		visual.SwapEvent(q, visual.Swap{I: i, J: j})
	}
}
