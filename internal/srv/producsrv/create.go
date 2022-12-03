package producsrv

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/lapitskyss/go_backend_2/internal/pkg/response"
)

func (s *ProductSrv) Create(ctx context.Context, product Product) (*Product, error) {
	err := product.validate()
	if err != nil {
		return nil, response.ErrBadRequest(err)
	}

	id := uuid.New()
	product.ID = id.String()

	err = s.es.SaveProduct(ctx, product)
	if err != nil {
		s.log.Error("Product service, method create, function SaveProduct", zap.Error(err))
		return nil, response.ErrInternal()
	}

	return &product, nil
}
