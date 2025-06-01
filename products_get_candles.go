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

type GetProductCandlesOptions struct {
	ProductID   string          `url:"-"`               // The trading pair (e.g. 'BTC-USD').
	Start       time.Time       `url:"start,unix"`      // The UNIX timestamp indicating the start of the time interval.
	End         time.Time       `url:"end,unix"`        // The UNIX timestamp indicating the end of the time interval.
	Granularity TimeGranularity `url:"granularity"`     // The timeframe each candle represents.
	Limit       *int            `url:"limit,omitempty"` // The number of candle buckets to be returned. By default, returns 350 (max 350).
}

type getProductCandlesResponse struct {
	Candles []Candles `json:"candles"`
}

// Get rates for a single product by product ID, grouped in buckets.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getcandles
func (s *ProductsService) GetProductCandles(
	ctx context.Context,
	opts *GetProductCandlesOptions,
) ([]Candles, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/products/%s/candles",
		s.client.baseURL,
		opts.ProductID,
	)

	var resp getProductCandlesResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		err = fmt.Errorf(
			"failed to fetch get product candles for product '%s': %w",
			opts.ProductID,
			err,
		)
	}

	return resp.Candles, err
}
