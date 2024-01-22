package coinbase

import (
	"context"
	"fmt"
	"net/http"
)

// Get gets information on a single product by product ID.
func (s *ProductsService) Get(ctx context.Context, id string) (*Product, error) {
	url := fmt.Sprintf("%s/api/v3/brokerage/products/%s", s.client.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get product HTTP request: %w", err)
	}

	var product Product
	err = s.client.do(req, http.StatusOK, &product)
	if err != nil {
		err = fmt.Errorf("failed to fetch product '%s': %w", id, err)
	}

	return &product, err
}
