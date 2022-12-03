package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"

	"github.com/lapitskyss/go_backend_2/internal/srv/producsrv"
)

func (s *Store) SaveProduct(ctx context.Context, product producsrv.Product) error {
	bdy, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("insert: marshall: %w", err)
	}

	req := esapi.CreateRequest{
		Index:      s.productAlias,
		DocumentID: product.ID,
		Body:       bytes.NewReader(bdy),
	}

	res, err := req.Do(ctx, s.es)
	if err != nil {
		return fmt.Errorf("save product: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("save product: response: %s", res.String())
	}

	return nil
}

type searchQueryResult struct {
	Hits searchHits `json:"hits"`
}

type searchHits struct {
	Hits []searchHitsHits `json:"hits"`
}

type searchHitsHits struct {
	Source producsrv.Product `json:"_source"`
}

func (s *Store) SearchProducts(ctx context.Context, query string) ([]producsrv.Product, error) {
	var reader io.Reader

	if query == "" {
		reader = strings.NewReader(searchAllProducts)
	} else {
		reader = strings.NewReader(fmt.Sprintf(searchProducts, query, query))
	}

	res, err := s.es.Search(
		s.es.Search.WithContext(ctx),
		s.es.Search.WithIndex(s.productIndex),
		s.es.Search.WithBody(reader),
	)
	if err != nil {
		return nil, fmt.Errorf("save product: request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search product: response: %s", res.String())
	}

	var qr searchQueryResult
	err = json.NewDecoder(res.Body).Decode(&qr)
	if err != nil {
		return nil, fmt.Errorf("search product: decode result: %w", err)
	}

	n := len(qr.Hits.Hits)
	products := make([]producsrv.Product, 0, n)
	for _, hit := range qr.Hits.Hits {
		products = append(products, hit.Source)
	}

	return products, nil
}

const searchAllProducts = `{
	"query": { "match_all" : {} },
	"size": 25,
	"from": 0
}`

const searchProducts = `{
	"query": {
		"dis_max":{
			"queries": [
				{
					"match": {
						"title": {
							"query": %q,
							"boost": 100
						}
					}
				},
				{
					"match": {
						"description": {
							"query": %q,
							"boost": 1
						}
					}
				}
			],
			"tie_breaker": 0,
			"boost": 1
		}
	},
	"size": 25,
	"from": 0
}`
