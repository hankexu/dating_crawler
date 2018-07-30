package engine

type Parser interface {
	Parse(contents []byte, user string) ParseResult
	Serialize() (name string, args interface{})
}

// Request 请求
type Request struct {
	URL    string
	Parser Parser
}

type ParserFunc func(content []byte, url string) ParseResult

// ParseResult 解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type NilParser struct{}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func CreateFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{parser: p, name: name}
}
