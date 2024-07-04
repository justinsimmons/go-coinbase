package coinbase

import (
	"context"
	"fmt"
)

// List Public Products.
// Get a list of the available currency pairs for trading.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproducts/
func (s *PublicService) ListProducts(ctx context.Context, options *ListProductsOptions) ([]Product, error) {
	u := s.client.baseURL + "/api/v3/brokerage/market/products"

	var productsResp listProductsResponse

	err := s.client.get(ctx, u, options, &productsResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of products: %w", err)
	}

	return productsResp.Products, err
}
