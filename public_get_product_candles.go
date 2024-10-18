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
