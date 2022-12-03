package producsrv

import (
	"context"

	"go.uber.org/zap"
)

type ProductSrv struct {
	es  ElasticStore
	log *zap.Logger
}

func InitProductSrv(es ElasticStore, log *zap.Logger) *ProductSrv {
	return &ProductSrv{
		es:  es,
		log: log,
	}
}

type ElasticStore interface {
	SaveProduct(ctx context.Context, product Product) error
	SearchProducts(ctx context.Context, query string) ([]Product, error)
}
