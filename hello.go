package hello

import "fmt"

type Hello struct{}

func (Hello) Print() string {
	s := "hello v1.0.0\n"
	fmt.Println(s)
	return s
}
