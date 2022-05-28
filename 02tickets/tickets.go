package tickets

import (
	"strconv"
)

type Tickets struct {
	count  int
	table  [][10]int
	sum    []int
	square []int
}

func (t Tickets) Run(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "error"
	}
	t.table = [][10]int{
		{1},
		{0, 1},
		{0, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}
	t.calc(n)
	return strconv.Itoa(t.count)
}

func (t *Tickets) calc(n int) {
	if n == 0 {
		return
	}
	t.sum = make([]int, len(t.table))
	t.square = make([]int, len(t.table))
	var cnt int
	for i := range t.table {
		cnt = 0
		for _, elem := range t.table[i] {
			cnt += elem
		}
		t.sum[i] = cnt
		t.square[i] = cnt * cnt
	}
	t.count = 0
	for _, elem := range t.square {
		t.count += elem
	}
	t.refillTable()
	t.calc(n - 1)
}

func (t *Tickets) refillTable() {
	t.table = make([][10]int, len(t.sum)+9)
	for i := 0; i < 10; i++ {
		for j := range t.sum {
			t.table[i+j][i] = t.sum[j]
		}
	}
}
