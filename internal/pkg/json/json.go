package json

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Parser struct {
	json jsoniter.API
}

func NewParser() *Parser {
	return &Parser{json: json}
}

func (p *Parser) Marshal(v interface{}) ([]byte, error) {
	return p.json.Marshal(v)
}

func (p *Parser) Unmarshal(data []byte, v interface{}) error {
	return p.json.Unmarshal(data, v)
}
