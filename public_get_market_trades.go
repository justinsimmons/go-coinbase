package coinbase

import (
	"context"
	"fmt"
)

// Get Public Market Trades.
// Get snapshot information by product ID about the last trades (ticks) and best bid/ask.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicmarkettrades/
func (s *PublicService) GetMarketTrades(ctx context.Context, options GetMarketTradeOptions) (*GetMarketTradesResponse, error) {
	u := fmt.Sprintf("%s/api/v3/brokerage/market/products/%s/ticker", s.client.baseURL, options.ProductID)

	var trades GetMarketTradesResponse
	err := s.client.get(ctx, u, &options, &trades)
	if err != nil {
		err = fmt.Errorf("failed to fetch public market trades for product '%s': %w", options.ProductID, err)
	}

	return &trades, err
}
