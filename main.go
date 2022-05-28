package main

import (
	"algorythms/02tickets"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

type ITask interface {
	Run(s string) string
}

func test(task ITask, path string) {
	started := time.Now()
	defer func() {
		fmt.Printf("Total time elapsed: %d ms", time.Since(started).Milliseconds())
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
		input = strings.TrimSpace(string(bIn)) //[:strings.IndexAny(string(bIn), "\n\r")]
		expected = strings.TrimSpace(string(bOut))
		res = task.Run(input)
		if res != expected {
			fmt.Printf("%s test.%d.in failed within %d ms\n", name, i, time.Since(instance).Milliseconds())
			return
		}
		fmt.Printf("%s test.%d.in success within %d ms\n", name, i, time.Since(instance).Milliseconds())
	}
	return
}

func main() {
	//sl := strlen.StringLength{}
	//test(sl, "strlen")
	tckt := tickets.Tickets{}
	test(tckt, "02tickets")
}
