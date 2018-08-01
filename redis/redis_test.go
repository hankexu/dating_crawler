package redis

import (
	r "github.com/gomodule/redigo/redis"
	"testing"
)

func TestConn_IsDuplicate(t *testing.T) {

	conn, err := r.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	v := Conn{conn}
	conn.Do("SET", "http://album.zhenai.com/u/1874886935", 1)
	exist := v.IsDuplicate("http://album.zhenai.com/u/1874886935")
	if !exist {
		t.Error("Url should be exist, but not")
	}

}
