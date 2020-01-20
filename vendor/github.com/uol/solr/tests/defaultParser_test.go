package tests

import (
	"testing"

	"github.com/uol/solr/parser"
	"github.com/uol/solr/solr"
)

func TestDefaultParser(t *testing.T) {
	settings := solr.SettingsSolrCore{}
	var params map[string]string
	instance := solr.New("http://localhost:8080", "produtos_digitais", false, false, settings, params, &parser.DefaultDocumentParser{})
	searchParams := &solr.SearchParams{}
	res, err := instance.Search(searchParams)
	if err != nil {
		t.Error(err.Error())
	}
	if res.Status != 0 {
		t.Error()
	}
	if res.Docs == "" {
		t.Error()
	}
}
