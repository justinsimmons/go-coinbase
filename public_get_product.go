package coinbase

import (
	"context"
	"fmt"
)

// Get Public Product.
// Get information on a single product by product ID.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproduct/
func (s *PublicService) GetProduct(ctx context.Context, id string) (*Product, error) {
	u := fmt.Sprintf("%s/api/v3/brokerage/market/products/%s", s.client.baseURL, id)

	var product Product
	err := s.client.get(ctx, u, nil, &product)
	if err != nil {
		err = fmt.Errorf("failed to fetch product '%s': %w", id, err)
	}

	return &product, err
}
