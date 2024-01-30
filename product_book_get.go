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

type getProductBookResponse struct {
	PriceBook PriceBook `json:"pricebook"`
}

type getProductBookOptions struct {
	ProductID string `url:"product_id"`
	Limit     *int   `url:"limit,omitempty"`
}

// GetProductBook gets a list of bids/asks for a single product.
// The amount of detail shown can be customized with the limit parameter.
func (s *ProductsService) GetProductBook(ctx context.Context, id string, limit *int) (*PriceBook, error) {
	options := getProductBookOptions{
		ProductID: id,
		Limit:     limit,
	}

	var productBook getProductBookResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/product_book", &options, &productBook)
	if err != nil {
		err = fmt.Errorf("failed to fetch product book for product '%s': %w", id, err)
	}

	return &productBook.PriceBook, err
}
