package solrts

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/uol/solr/solr"
)

const (
	rawDoc      string = "docs"
	rawResponse string = "response"
	rawID       string = "id"
	rawMetric   string = "metric"
	rawType     string = "type"
	rawChildDoc string = "_childDocuments_"
	rawTagKey   string = "tag_key"
	rawTagValue string = "tag_value"
)

// TSDocumentParser - a ts document parser
type TSDocumentParser struct {
}

// Doc - a ts document structs
type Doc struct {
	ID     string
	Metric string
	Type   string
	Tags   []Tag
}

// Tag - a ts tag struct
type Tag struct {
	Name  string
	Value string
}

// Parse - parses the pure document input from JSON
func (p *TSDocumentParser) Parse(raw []byte) (interface{}, error) {
	res := &solr.Response{}
	var docsArray []Doc

	_, err := jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		doc := Doc{}
		if doc.ID, err = jsonparser.GetString(value, rawID); err != nil {
			return
		}

		if doc.Metric, err = jsonparser.GetString(value, rawMetric); err != nil {
			return
		}

		if doc.Type, err = jsonparser.GetString(value, rawType); err != nil {
			return
		}

		child, _, _, err := jsonparser.Get(value, rawChildDoc)
		if err != nil {
			return
		}
		if doc.Tags, err = parserChildsDocumentsObject(child); err != nil {
			return
		}

		docsArray = append(docsArray, doc)
	}, rawResponse, rawDoc)
	if err != nil {
		return res.Docs, err
	}
	fmt.Println(res.Docs)
	res.Docs = docsArray
	return res.Docs, nil
}

// parserChildsDocumentsObject - parses a ts child document from JSON
func parserChildsDocumentsObject(raw []byte) ([]Tag, error) {
	var tagsArray []Tag
	_, err := jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tag := Tag{}
		if tag.Name, err = jsonparser.GetString(value, rawTagKey); err != nil {
			return
		}
		fmt.Print(tag.Name)
		if tag.Value, err = jsonparser.GetString(value, rawTagValue); err != nil {
			return
		}
		tagsArray = append(tagsArray, tag)
	})
	if err != nil {
		return nil, err
	}
	return tagsArray, nil
}
