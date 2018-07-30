package parser

import (
	"crawler/config"
	"crawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			URL:    url,
			Parser: NewProfileParser(string(m[2])),
		})

	}
	cityMatches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityMatches {
		result.Requests = append(result.Requests, engine.Request{
			URL:    string(m[1]),
			Parser: engine.CreateFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
