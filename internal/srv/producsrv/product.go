package producsrv

import "errors"

type Product struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint64 `json:"price"`
	Quantity    uint64 `json:"quantity"`
}

func (p *Product) validate() error {
	if p.Title == "" {
		return errors.New("product title is required")
	}
	if len(p.Title) > 255 {
		return errors.New("product title is to long")
	}

	if p.Description == "" {
		return errors.New("product description is required")
	}
	if len(p.Description) > 10000 {
		return errors.New("product description is to long")
	}

	return nil
}
