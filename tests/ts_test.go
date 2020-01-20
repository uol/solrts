package tests

import (
	"fmt"
	"testing"

	lib "github.com/uol/solr/solr"
	"github.com/uol/solrts"
)

func BenchmarkSolrLib(b *testing.B) {
	settings := lib.SettingsSolrCore{}
	var params map[string]string
	instance := lib.New("http://localhost:8080", "produtos_digitais", false, false, settings, params, &solrts.TSDocumentParser{})
	searchParams := &lib.SearchParams{}
	res, err := instance.Search(searchParams)
	if err != nil {
		b.Error(err.Error())
	}
	fmt.Printf("%+v", res)
}
