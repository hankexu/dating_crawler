package parser_test

import (
	"fmt"
	"testing"
)

func a() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Test(t *testing.T) {
	f := a()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
