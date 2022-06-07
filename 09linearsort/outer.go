package _9linearsort

func OuterSort(a []int) []int {
	var b, c []int
	for {
		b, c = split(a)
		if b == nil {
			return c
		}
		if c == nil {
			return b
		}
		a = merge(b, c)
	}
}

func split(a []int) ([]int, []int) {
	if len(a) == 0 {
		return nil, nil
	}
	var b, c []int
	var current = &b
	b = append(b, a[0])
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			if current == &b {
				current = &c
			} else {
				current = &b
			}
		}
		*current = append(*current, a[i])
	}
	return b, c
}

func merge(b, c []int) []int {
	var i, j int
	a := make([]int, 0, len(b)+len(c))
	var first = true
	if len(b) == 0 {
		return c
	}
	if len(c) == 0 {
		return b
	}
	if b[0] >= c[0] {
		a = append(a, c[0])
		j++
	} else {
		a = append(a, b[0])
		i++
	}
	appenda := func() {
		if first {
			a = append(a, b[i])
			i++
		} else {
			a = append(a, c[j])
			j++
		}
	}
	for {
		if i == len(b) {
			a = append(a, c[j:]...)
			return a
		}
		if j == len(c) {
			a = append(a, b[i:]...)
			return a
		}
		switch {
		case a[len(a)-1] > b[i] && b[i] > c[j]:
			first = !first
			appenda()
		case a[len(a)-1] > b[i] && b[i] <= c[j]:
			appenda()
		case a[len(a)-1] <= b[i] && b[i] <= c[j]:
			if !first {
				first = !first
			}
			appenda()
		case a[len(a)-1] <= c[j] && c[j] <= b[i]:
			if first {
				first = !first
			}
			appenda()
		case a[len(a)-1] > b[i] && a[len(a)-1] <= c[j]:
			if first {
				first = !first
			}
			appenda()
		case a[len(a)-1] <= b[i] && a[len(a)-1] > c[j]:
			if !first {
				first = !first
			}
			appenda()
		}
	}
}
