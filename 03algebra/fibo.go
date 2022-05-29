package _3algebra

func FiboIter(n int) int {
	if n == 0 {
		return 0
	}
	f1, f2 := 0, 1
	for i := 1; i < n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
