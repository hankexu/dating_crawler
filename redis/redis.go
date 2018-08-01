package redis

import (
	"crawler/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"regexp"
)

type Conn struct {
	redis.Conn
}

func CreateConn(port int) (Conn, error) {
	var conn Conn
	c, err := redis.Dial("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return conn, err
	}
	conn.Conn = c
	return conn, nil
}

var profileRe = regexp.MustCompile(`http://album.zhenai.com/u/[0-9]+`)

func (c Conn) IsDuplicate(url string) bool {
	if !profileRe.Match([]byte(url)) {
		return false
	}
	if url == config.EntryUrl {
		return false
	}
	exists, _ := redis.Bool(c.Do("EXISTS", url))
	if exists {
		return true
	} else {
		c.Do("SET", url, 1)
	}
	return false
}
