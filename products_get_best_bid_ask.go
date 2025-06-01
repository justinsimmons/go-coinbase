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

type GetBestBidAskResponse struct {
	PriceBooks []PriceBook `json:"pricebooks"`
}

type GetBestBidAskOptions struct {
	ProductIDs []string `url:"product_ids,omitempty"`
}

// GetBestBidAsk gets the best bid/ask for all products. A subset of all
// products can be returned instead by using the product_ids input.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getbestbidask
func (s *ProductsService) GetBestBidAsk(
	ctx context.Context,
	opts *GetBestBidAskOptions,
) (*GetBestBidAskResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/best_bid_ask"

	var resp GetBestBidAskResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		err = fmt.Errorf(
			"failed to fetch best bid/ask for products '%v': %w",
			ids,
			err,
		)
	}

	return &resp, err
}
