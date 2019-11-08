package hello

import "fmt"

type Hello struct{}

func (Hello) Print() string {
	s := "hello v0.0.1\n"
	fmt.Println(s)
	return s
}
