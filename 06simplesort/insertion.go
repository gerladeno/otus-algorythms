package _6simplesort

import "algorythms/sortingcommon"

type Insertion struct {
	array []int
}

func (in *Insertion) GetSlice() []int {
	return in.array
}

func (in *Insertion) Run(s string) string {
	var err error
	in.array, err = sortingcommon.PrepareSlice(&s)
	if err != nil {
		return err.Error()
	}
	in.sort()
	result := sortingcommon.ReturnString(in.array)
	return *result
}

func (in *Insertion) sort() {
	var idx, tmp int
	for i := 0; i < len(in.array)-1; i++ {
		if in.array[i+1] < in.array[i] {
			idx = sortingcommon.BinarySearch(in, 0, i, in.array[i+1])
			tmp = in.array[i+1]
			for j := i; j >= idx; j-- {
				in.array[j+1] = in.array[j]
			}
			in.array[idx] = tmp
		}
	}
}
