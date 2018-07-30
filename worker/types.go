package worker

import (
	"crawler/config"
	"crawler/engine"
	"crawler/zhenai/parser"
	"fmt"
	"log"
)

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializeParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.URL,
		Parser: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, request := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(request))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	_parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		URL:    r.Url,
		Parser: _parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing request %v:%v\n", req, err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func deserializeParser(p SerializeParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.CreateFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.CreateFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				userName), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"arg: %v", p.Args)
		}
	case config.NilParse:
		return engine.NilParser{}, nil
	default:
		return nil, fmt.Errorf("unknown parser name:%v", p.Name)
	}
}
