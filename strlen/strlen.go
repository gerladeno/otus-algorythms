package strlen

import "strconv"

type StringLength struct {
}

func (sl StringLength) Run(s string) string {
	return strconv.Itoa(len(s))
}
