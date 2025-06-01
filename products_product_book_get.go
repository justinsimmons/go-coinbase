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

type GetProductBookResponse struct {
	PriceBook      PriceBook `json:"pricebook"`
	Last           *string   `json:"last,omitempty"`
	MidMarket      *string   `json:"mid_market,omitempty"`
	SpreadBPS      *string   `json:"spread_bps,omitempty"`
	SpreadAbsolute *string   `json:"spread_absolute,omitempty"`
}

type GetProductBookOptions struct {
	ProductID                 string  `url:"product_id"`                            // The trading pair (e.g. 'BTC-USD').
	Limit                     *int    `url:"limit,omitempty"`                       // The number of bid/asks to be returned.
	AggregationPriceIncrement *string `url:"aggregation_price_increment,omitempty"` // The minimum price intervals at which buy and sell orders are grouped or combined in the order book.
}

// Get a list of bids/asks for a single product. The amount of detail shown
// can be customized with the limit parameter.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getproductbook
func (s *ProductsService) GetProductBook(
	ctx context.Context,
	opts *GetProductBookOptions,
) (*GetProductBookResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/product_book"

	var resp GetProductBookResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		err = fmt.Errorf(
			"failed to fetch product book for product '%s': %w",
			opts.ProductID,
			err,
		)
	}

	return &resp, err
}
