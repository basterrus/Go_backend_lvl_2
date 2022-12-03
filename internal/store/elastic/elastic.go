package elastic

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type Store struct {
	es           *elasticsearch.Client
	productIndex string
	productAlias string
}

func InitElasticsearch(addresses []string) (*Store, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	store := &Store{
		es:           es,
		productIndex: "product",
		productAlias: "product_alias",
	}

	err = store.createProductIndex()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) createProductIndex() error {
	res, err := s.es.Indices.Exists([]string{s.productIndex})
	if err != nil {
		return fmt.Errorf("cannot check index existence: %w", err)
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return fmt.Errorf("error in index existence response: %s", res.String())
	}

	body := `{
		"settings": {
			"analysis": {
				"filter": {
					"ngram_filter": {
						"type": "edge_ngram",
						"min_gram": 3,
						"max_gram": 20
					}
				},
				"analyzer": {
					"ngram_analyzer": {
						"type": "custom",
						"tokenizer": "standard",
						"filter": [
							"lowercase",
							"ngram_filter"
						]
					}
				}
			}
		},
		"mappings": {
			"properties": {
				"title": { "type": "text", "analyzer": "ngram_analyzer" },
				"description": { "type": "text", "analyzer": "ngram_analyzer" }
			}
		}
	}`

	res, err = s.es.Indices.Create(s.productIndex,
		s.es.Indices.Create.WithBody(strings.NewReader(body)),
	)
	if err != nil {
		return fmt.Errorf("cannot create index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index creation response: %s", res.String())
	}

	res, err = s.es.Indices.PutAlias([]string{s.productIndex}, s.productAlias)
	if err != nil {
		return fmt.Errorf("cannot create index alias: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in index alias creation response: %s", res.String())
	}

	return nil
}
