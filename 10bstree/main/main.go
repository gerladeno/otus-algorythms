package main

import (
	_0bstree "algorythms/10bstree"
	"fmt"
	"math/rand"
	"time"
)

//func main() {
//	a := _0bstree.AVLTree{Key: 10}
//	a.Insert(5, nil)
//	a.Insert(10, nil)
//	a.Insert(12, nil)
//	a.Insert(7, nil)
//	a.Insert(3, nil)
//	a.Insert(11, nil)
//	a.Insert(20, nil)
//	a.Insert(17, nil)
//	a.Insert(25, nil)
//	a.Draw()
//	fmt.Printf("%#v\n", a.Search(12))
//	fmt.Printf("%#v\n", a.Search(13))
//	a.Delete(25)
//	a.Draw()
//	a.Delete(20)
//	a.Draw()
//	a.Delete(12)
//	a.Draw()
//	a.Delete(5)
//	a.Draw()
//}

const N = 100_000

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := _0bstree.AVLTree{Key: r.Int()}
	b := _0bstree.AVLTree{Key: 0}
	started := time.Now()
	for i := 1; i < N; i++ {
		a.Insert(int(r.Int31()), nil)
		b.Insert(i, nil)
		if i%(N/100) == 0 {
			fmt.Printf("\r%d%%", 100*i/N)
		}
	}
	fmt.Printf("\ntook %d ms to fill with %d elems for sorted\n", time.Since(started).Milliseconds(), N)
	started = time.Now()
	for i := 0; i < N/10; i++ {
		a.Delete(int(r.Int31()))
		if i%(N/1000) == 0 {
			fmt.Printf("\r%d%%", 1000*i/N)
		}
	}
	fmt.Printf("\ntook %d ms to delete %d elems from random\n", time.Since(started).Milliseconds(), N/10)
	started = time.Now()
	for i := 0; i < N/10; i++ {
		b.Delete(int(r.Int31()))
		if i%(N/1000) == 0 {
			fmt.Printf("\r%d%%", 1000*i/N)
		}
	}
	fmt.Printf("\ntook %d ms to delete %d elems from sorted\n", time.Since(started).Milliseconds(), N/10)
}
