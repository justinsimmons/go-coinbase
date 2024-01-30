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

type ExpiringContractStatus string

const (
	ExpiringContractStatusUnknown   ExpiringContractStatus = "UNKNOWN_EXPIRING_CONTRACT_STATUS"
	ExpiringContractStatusUnexpired ExpiringContractStatus = "STATUS_UNEXPIRED"
	ExpiringContractStatusExpired   ExpiringContractStatus = "STATUS_EXPIRED"
	ExpiringContractStatusAll       ExpiringContractStatus = "STATUS_ALL"
)

type ListProductsOptions struct {
	Limit                  *int32                  `url:"limit,omitempty"`        // A limit describing how many products to return.
	Offset                 *int32                  `url:"limit,omitempty"`        // Number of products to offset before returning.
	ProductType            *ProductType            `url:"product_type,omitempty"` // Type of products to return.
	ProductIDs             []string                `url:"product_ids,omitempty"`  // List of product IDs to return.
	ContractExpiryType     *ContractExpiryType     `url:"contract_expiry_type,omitempty"`
	ExpiringContractStatus *ExpiringContractStatus `url:"expiring_contract_status,omitempty"`
}

type listProductsResponse struct {
	Products       []Product `json:"products"`     // Array of objects, each representing one product.
	NumberProducts int32     `json:"num_products"` // Number of products that were returned.
}

func (s *ProductsService) List(ctx context.Context, options *ListProductsOptions) ([]Product, error) {
	var productsResp listProductsResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/products", options, &productsResp)
	if err != nil {
		err = fmt.Errorf("failed to fetch list of products: %w", err)
	}

	return productsResp.Products, err
}
