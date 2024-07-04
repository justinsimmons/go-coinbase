package coinbase

import (
	"context"
	"fmt"
)

// Get Public Product Candles.
// Get rates for a single product by product ID, grouped in buckets.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpubliccandles/
func (s *PublicService) GetProductCandles(ctx context.Context, options GetProductCandlesOptions) ([]Candles, error) {
	u := fmt.Sprintf("%s/api/v3/brokerage/market/products/%s/candles", s.client.baseURL, options.ProductID)

	var candlesResp getProductCandlesResponse
	err := s.client.get(ctx, u, &options, &candlesResp)
	if err != nil {
		err = fmt.Errorf("failed to fetch get product candles for product '%s': %w", options.ProductID, err)
	}

	return candlesResp.Candles, err
}
