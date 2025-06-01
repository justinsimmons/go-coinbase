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

type ListProductsOptions struct {
	Limit                  *int32                  `url:"limit,omitempty"`        // A limit describing how many products to return.
	Offset                 *int32                  `url:"limit,omitempty"`        // Number of products to offset before returning.
	ProductType            *ProductType            `url:"product_type,omitempty"` // Type of products to return.
	ProductIDs             []string                `url:"product_ids,omitempty"`  // List of product IDs to return.
	ContractExpiryType     *ContractExpiryType     `url:"contract_expiry_type,omitempty"`
	ExpiringContractStatus *ExpiringContractStatus `url:"expiring_contract_status,omitempty"` // Status of expiring contract products to return. Default is UNEXPIRED.
}

type listProductsResponse struct {
	Products       []Product `json:"products"`     // Array of objects, each representing one product.
	NumberProducts int32     `json:"num_products"` // Number of products that were returned.
}

// Get a list of the available currency pairs for trading.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getproducts
func (s *ProductsService) List(
	ctx context.Context,
	opts *ListProductsOptions,
) ([]Product, error) {

	u := s.client.baseURL + "/api/v3/brokerage/products"

	var resp listProductsResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		err = fmt.Errorf("failed to fetch list of products: %w", err)
	}

	return resp.Products, err
}
