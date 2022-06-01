package _7pyramidsort

import (
	"algorythms/sortingcommon"
	"algorythms/visual"
)

type Heap struct {
	Ch chan struct{}
	a  []int
}

func (h *Heap) Run(s string) string {
	var err error
	h.a, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	h.sort()
	result := sortingcommon.ReturnString(h.a)
	return *result
}

func (h *Heap) GetSlice() []int {
	return h.a
}

func (h *Heap) sort() {
	for i := len(h.a)/2 - 1; i >= 0; i-- {
		h.heapify(i, len(h.a))
	}
	for i := len(h.a) - 1; i >= 0; i-- {
		visual.SwapEvent(h, visual.Swap{I: 0, J: i})
		h.heapify(0, i)
	}
}

func (h *Heap) heapify(root, size int) {
	l := 2*root + 1
	r := l + 1
	x := root
	<-h.Ch
	if l < size && h.a[x] < h.a[l] {
		x = l
	}
	if r < size && h.a[x] < h.a[r] {
		x = r
	}
	if x == root {
		return
	}
	visual.SwapEvent(h, visual.Swap{I: x, J: root})
	//sortingcommon.Swap(h, x, l)
	h.heapify(x, size)
}

func (h *Heap) sortXLR(x, l, r int) {
	<-h.Ch
	max := x
	if h.a[x] < h.a[l] {
		max = l
	}
	if h.a[max] < h.a[r] {
		max = r
	}
	if max != x {
		visual.SwapEvent(h, visual.Swap{I: x, J: max})
		//sortingcommon.Swap(h, x, l)
	}
}
