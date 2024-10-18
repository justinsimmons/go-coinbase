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
