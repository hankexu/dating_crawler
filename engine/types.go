package engine

// Request 请求
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

type ParserFunc func(content []byte, url string) ParseResult

// ParseResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// NilParser 空解析器
func NilParser([]byte, string) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
