package parser

import (
	"crawler/config"
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市列表解析器
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	match := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		result.Requests = append(result.Requests, engine.Request{
			URL:    string(m[1]),
			Parser: engine.CreateFuncParser(ParseCity, config.ParseCity),
		})

	}

	return result
}
