package main

import (
	_8quicksort "algorythms/08quicksort"
	"algorythms/visual"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

const maxStringLen = 200

type ITask interface {
	Run(s string) string
}

func test(task ITask, path string) {
	started := time.Now()
	defer func() {
		fmt.Printf("Total time elapsed: %d ns", time.Since(started).Nanoseconds())
	}()
	var instance time.Time
	name := reflect.TypeOf(task).Name()
	var err error
	var bIn, bOut []byte
	var res, input, expected string
	for i := 0; ; i++ {
		instance = time.Now()
		in := fmt.Sprintf("%s/test.%d.in", path, i)
		bIn, err = os.ReadFile(in)
		if os.IsNotExist(err) {
			break
		}
		out := fmt.Sprintf("%s/test.%d.out", path, i)
		bOut, err = os.ReadFile(out)
		if os.IsNotExist(err) {
			break
		}
		input = strings.TrimSpace(string(bIn))
		expected = strings.TrimSpace(string(bOut))
		res = task.Run(input)
		if len(expected) > maxStringLen {
			expected = expected[:maxStringLen]
		}
		if len(res) > maxStringLen {
			res = res[:maxStringLen]
		}
		if res != expected {
			fmt.Printf("%s test.%d.in failed within %d ns.\nExpected: %s\n     got: %s\n", name, i, time.Since(instance).Nanoseconds(), expected, res)
			continue
		}
		fmt.Printf("%s test.%d.in success within %d ns\n", name, i, time.Since(instance).Nanoseconds())
	}
	return
}

func main() {
	ch := make(chan struct{})
	item := &_8quicksort.Quick{Ch: ch}
	go func() {
		test(item, "sorting-tests/0.random")
	}()
	if err := visual.Run(ch, item); err != nil {
		fmt.Println(err.Error())
	}
}

//func main() {
//	item := &_7pyramidsort.Selection{}
//	test(item, "sorting-tests/0.random")
//}
