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

// GetBestBidAsk gets the best bid/ask for all products. A subset of all products can be
// returned instead by using the product_ids input.
func (s *ProductsService) GetBestBidAsk(ctx context.Context, ids ...string) (*GetBestBidAskResponse, error) {
	options := struct {
		ProductIDs []string `url:"product_ids"`
	}{
		ProductIDs: ids,
	}

	var bidAsk GetBestBidAskResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/best_bid_ask", &options, &bidAsk)
	if err != nil {
		err = fmt.Errorf("failed to fetch best bid/ask for products '%v': %w", ids, err)
	}

	return &bidAsk, err
}
