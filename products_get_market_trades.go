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
	"time"
)

type GetMarketTradeOptions struct {
	ProductID string     `url:"-"`                    // The trading pair, i.e., 'BTC-USD'.
	Limit     int        `url:"limit"`                // Number of trades to return.
	Start     *time.Time `url:"start,unix,omitempty"` // The UNIX timestamp indicating the start of the time interval.
	End       *time.Time `url:"end,unix,omitempty"`   // The UNIX timestamp indicating the end of the time interval.
}

type GetMarketTradesResponse struct {
	Trades  []Trade `json:"trades"`
	BestBid *string `json:"best_bid,omitempty"` // The best bid for the `product_id`, in quote currency.
	BestAsk *string `json:"best_ask,omitempty"` // The best ask for the `product_id`, in quote currency.
}

// Get snapshot information by product ID about the last trades (ticks) and
// best bid/ask.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getmarkettrades
func (s *ProductsService) GetMarketTrades(
	ctx context.Context,
	opts *GetMarketTradeOptions,
) (*GetMarketTradesResponse, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/products/%s/ticker",
		s.client.baseURL,
		opts.ProductID,
	)

	var resp GetMarketTradesResponse

	err := s.client.get(ctx, u, &opts, &resp)
	if err != nil {
		err = fmt.Errorf(
			"failed to fetch get market trades for product '%s': %w",
			opts.ProductID,
			err,
		)
	}

	return &resp, err
}
