package parser

import (
	"encoding/json"
)

type responseRaw struct {
	Data dataRaw `json:"response"`
}

type dataRaw struct {
	NumFound  int           `json:"numFound"`
	Start     int           `json:"start"`
	Documents []documentRaw `json:"docs"`
}
type documentRaw map[string]interface{}

// DefaultDocumentParser - Default parser
type DefaultDocumentParser struct {
}

// Parse - parses a document and return a default struct
func (p *DefaultDocumentParser) Parse(raw []byte) (interface{}, error) {
	var resp responseRaw
	err := json.Unmarshal(raw, &resp)

	if err != nil {
		return responseRaw{}, err
	}
	return resp.Data.Documents, err
}
