package coinbase

import (
	"context"
	"fmt"
)

// Get Public Product Book.
// Get a list of bids/asks for a single product.
// The amount of detail shown can be customized with the limit parameter.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproductbook/
func (s *PublicService) GetProductBook(ctx context.Context, opts GetProductBookOptions) (*PriceBook, error) {
	u := s.client.baseURL + "/api/v3/brokerage/market/product_book"

	var productBook getProductBookResponse
	err := s.client.get(ctx, u, &opts, &productBook)
	if err != nil {
		err = fmt.Errorf("failed to fetch product book for product '%s': %w", opts.ProductID, err)
	}

	return &productBook.PriceBook, err
}
