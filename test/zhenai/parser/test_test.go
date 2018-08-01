package parser_test

import (
	"fmt"
	"regexp"
	"testing"
)

func a() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

var profileRe = regexp.MustCompile(`http://album.zhenai.com/u/[0-9]+`)

func Test(t *testing.T) {
	fmt.Println(profileRe.Match([]byte("http://album.zhenai.com/u/110171680")))

}
