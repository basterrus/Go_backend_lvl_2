package producsrv

import (
	"context"

	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/pkg/response"
)

func (s *ProductSrv) Search(ctx context.Context, query string) ([]Product, error) {
	products, err := s.es.SearchProducts(ctx, query)
	if err != nil {
		s.log.Error("Product service, method search, function SearchProducts", zap.Error(err))
		return nil, response.ErrInternal()
	}

	return products, nil
}
