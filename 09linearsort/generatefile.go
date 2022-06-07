package _9linearsort

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	filename = `random16bit.bin`
	elements = 100_000_000
)

func Generate() error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Sync()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < elements; i++ {
		if i%(elements/100) == 0 {
			fmt.Printf("\r%d%%", 100*i/elements)
		}
		if err = binary.Write(file, binary.BigEndian, uint16(r.Int31())); err != nil {
			return err
		}
	}
	return file.Close()
}

func ReadBinary() ([]int, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	result := make([]int, 0, elements)
	var tmp uint16
	for {
		err = binary.Read(file, binary.BigEndian, &tmp)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		result = append(result, int(tmp))
	}
	return result, nil
}
