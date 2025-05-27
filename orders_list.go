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

type ListOrdersOptions struct {
	OrderIDs             []string              `url:"order_ids,omitempty"`              // ID(s) of order(s).
	ProductIDs           []string              `url:"product_ids,omitempty"`            // Optional string of the product ID(s). Defaults to null, or fetch for all products.
	ProductType          *ProductType          `url:"product_type,omitempty"`           // Returns orders matching this product type. By default, returns all product types.
	OrderStatuses        []OrderStatus         `url:"order_status,omitempty"`           // Only returns orders matching the specified order statuses.
	TimeInForces         TimeInForce           `url:"time_in_forces,omitempty"`         // Only orders matching this time in force(s) are returned. Default is to return all time in forces.
	OrderTypes           []OrderType           `url:"order_types,omitempty"`            // Only returns orders matching the specified order types (e.g. MARKET). By default, returns all order types.
	OrderSide            *Side                 `url:"order_side,omitempty"`             // Only returns the orders matching the specified side (e.g. 'BUY', 'SELL'). By default, returns all sides.
	StartDate            *time.Time            `url:"start_date,omitempty"`             // The start date to fetch orders from (inclusive). If provided, only orders created after this date will be returned.
	EndDate              *time.Time            `url:"end_date,omitempty"`               // The end date to fetch orders from (exclusive). If provided, only orders with creation time before this date will be returned.
	OrderPlacementSource *OrderPlacementSource `url:"order_placement_source,omitempty"` // Only returns the orders matching this placement source. By default, returns RETAIL_ADVANCED placement source.
	ContractExpiryType   *ContractExpiryType   `url:"contract_expiry_type,omitempty"`   // Only returns the orders matching the contract expiry type. Only applicable if product_type is set to FUTURE.
	AssetFilters         []string              `url:"asset_filters,omitempty"`          // Only returns the orders where the quote, base or underlying asset matches the provided asset filter(s) (e.g. 'BTC').
	RetailPortfolioID    *string               `url:"retail_portfolio_id,omitempty"`    // (Deprecated) Only orders matching this retail portfolio id are returned. Only applicable for legacy keys. CDP keys will default to the key's permissioned portfolio.
	Limit                *int32                `url:"limit,omitempty"`                  // The number of orders to display per page (no default amount). If has_next is true, additional pages of orders are available to be fetched. Use the cursor parameter to start on a specified page.
	Cursor               *string               `url:"cursor,omitempty"`                 // For paginated responses, returns all responses that come after this value.
	SortBy               *string               `url:"sort_by,omitempty"`                // Sort results by a field, results use unstable pagination. Default is to sort by creation time.
	UserNativeCurrency   *string               `url:"user_native_currency,omitempty"`   // (Deprecated) Native currency to fetch order with. Default is USD.
}

type ListOrdersResponse struct {
	Orders []Order `json:"orders"` // A list of orders matching the query.
	PaginatedResponse
	Sequence *int64 `json:"sequence,omitempty"` // **(Deprecated)** The sequence of the db at which this state was read.
}

// Get a list of orders filtered by optional query parameters (product_id, order_status, etc).
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_gethistoricalorders
func (s *OrdersService) List(
	ctx context.Context,
	opts *ListOrdersOptions,
) (*ListOrdersResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/historical/batch"

	var resp ListOrdersResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}

	return &resp, err
}
