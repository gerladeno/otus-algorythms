package _3algebra

import (
	"math"
	"math/big"
	"strconv"
)

var phi = (1 + math.Sqrt(5)) / 2.0

type Fibo struct {
}

func (f Fibo) Run(s string) string {
	n, err := strconv.Atoi(s)
	if err != nil {
		return "error"
	}
	result := FiboGoldenSection(n)
	return result.String()
}

func FiboIter(n int) *big.Int {
	if n == 0 {
		return &big.Int{}
	}
	var f1, f2, tmp big.Int
	f1.SetInt64(0)
	f2.SetInt64(1)
	for i := 1; i < n; i++ {
		tmp.Set(&f2)
		f2.Add(&f1, &f2)
		f1.Set(&tmp)
		//fmt.Printf("%s %s\n", f1.String(), f2.String())
	}
	return &f2
}

func FiboGoldenSection(n int) *big.Int {
	if n == 0 {
		return &big.Int{}
	}
	bigPhi := big.NewFloat(phi)
	//bigPhi := Zero().Sqrt(big.NewFloat(5))
	//bigPhi.Add(bigPhi, big.NewFloat(1))
	//bigPhi.Mul(bigPhi, big.NewFloat(0.5))
	result, _ := Zero().Add(Zero().Mul(Pow(bigPhi, n), big.NewFloat(1/math.Sqrt(5))), big.NewFloat(0.5)).Int(nil)
	return result
}

func Pow(val *big.Float, power int) *big.Float {
	result := big.NewFloat(1.0)
	for power >= 1 {
		if power%2 == 1 {
			result.Mul(result, val)
		}
		power /= 2
		val.Mul(val, val)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(big.MaxPrec)
	return r
}
