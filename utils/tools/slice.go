package tools

import "fmt"

type Slice struct{}

func (s *Slice) contains(slice []string, str string) bool {
	fmt.Println(str)
	fmt.Println(s)
	for _, a := range slice {
		if a == str {
			return true
		}
	}
	return false
}
