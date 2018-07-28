package fetcher

import (
	"fmt"
	"testing"
)

func TestFetcher(t *testing.T) {
	contents, err := Fetch("http://album.zhenai.com/u/1874886935")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
}
