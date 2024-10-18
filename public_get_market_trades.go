// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

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
