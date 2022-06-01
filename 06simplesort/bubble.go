package _6simplesort

import (
	"algorythms/sortingcommon"
	"algorythms/visual"
)

type Bubble struct {
	Ch    chan struct{}
	array []int
}

func (b *Bubble) Run(s string) string {
	var err error
	b.array, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	b.sort()
	result := sortingcommon.ReturnString(b.array)
	return *result
}

func (b *Bubble) GetSlice() []int {
	return b.array
}

func (b *Bubble) sort() {
	var swapped bool
	for i := range b.array {
		swapped = false
		for j := 0; j < len(b.array)-1-i; j++ {
			<-b.Ch
			if b.array[j] > b.array[j+1] {
				visual.SwapEvent(b, visual.Swap{I: j, J: j + 1})
				//sortingcommon.Swap(b, j, j+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
