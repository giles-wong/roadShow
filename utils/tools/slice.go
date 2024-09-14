package tools

import "fmt"

type Slice struct {
	Slice []string
}

func (s *Slice) SliceContains(str string) bool {
	fmt.Println(str)
	fmt.Println(s)
	for _, a := range s.Slice {
		if a == str {
			return true
		}
	}
	return false
}
