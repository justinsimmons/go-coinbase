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

// List Public Products.
// Get a list of the available currency pairs for trading.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproducts/
func (s *PublicService) ListProducts(ctx context.Context, options *ListProductsOptions) ([]Product, error) {
	u := s.client.baseURL + "/api/v3/brokerage/market/products"

	var productsResp listProductsResponse

	err := s.client.get(ctx, u, options, &productsResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of products: %w", err)
	}

	return productsResp.Products, err
}
