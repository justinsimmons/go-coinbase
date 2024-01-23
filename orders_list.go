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
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type OrderSide string

const (
	OrderSideUnknown OrderSide = "UNKNOWN_ORDER_SIDE"
	OrderSideBuy     OrderSide = "BUY"
	OrderSideSell    OrderSide = "SELL"
)

type ListOrdersOptions struct {
	ProductID            *string               `url:"product_id,omitempty"`             // Optional string of the product ID. Defaults to null, or fetch for all products.
	OrderStatus          []string              `url:"order_status,omitempty"`           // A list of order statuses.
	Limit                *int32                `url:"limit,omitempty"`                  // Pagination limit with a default of 1000 (which is also the maximum). If has_next is true, additional orders are available to be fetched with pagination; also the cursor value in the response can be passed as cursor parameter in the subsequent request.
	StartDate            *time.Time            `url:"start_date,omitempty"`             // Start date to fetch orders from, inclusive.
	EndDate              *time.Time            `url:"end_date,omitempty"`               // An optional end date for the query window, exclusive. If provided only orders with creation time before this date will be returned.
	OrderType            *OrderType            `url:"order_type,omitempty"`             // Type of orders to return. Default is to return all order types.
	OrderSide            *OrderSide            `url:"order_side,omitempty"`             // Only orders matching this side are returned. Default is to return all sides.
	Cursor               *string               `url:"cursor,omitempty"`                 // Cursor used for pagination. When provided, the response returns responses after this cursor.
	ProductType          *ProductType          `url:"product_type,omitempty"`           // Only orders matching this product type are returned. Default is to return all product types.
	OrderPlacementSource *OrderPlacementSource `url:"order_placement_source,omitempty"` // Only orders matching this placement source are returned. Default is to return RETAIL_ADVANCED placement source.
	ContractExpiryType   *ContractExpiryType   `url:"contract_expiry_type,omitempty"`   // Only orders matching this contract expiry type are returned. Filter is only applied if product_type is set to FUTURE in the request.
	AssetFilters         []string              `url:"asset_filters,omitempty"`          // Only returns orders where the quote, base, or underlying asset matches the provided asset filter(s), i.e., 'BTC'.
	RetailPortfolioID    *string               `url:"retail_portfolio_id,omitempty"`    // Only orders matching this retail portfolio id are returned. Default is to return orders for all retail portfolio ids.
}

type ListOrdersResponse struct {
	Orders   []Order `json:"orders"`   // A list of orders matching the query.
	Sequence *string `json:"sequence"` // The sequence of the db at which this state was read.
	HasNext  bool    `json:"has_next"` // Whether there are additional pages for this query.
	Cursor   *string `json:"cursor"`   // Cursor for paginating. Users can use this string to pass in the next call to this endpoint, and repeat this process to fetch all fills through pagination.
}

// List gets a list of orders filtered by optional query parameters (product_id, order_status, etc).
//
//  1. The maximum number of OPEN orders returned is 1000. If you have more than 1000 open, its recommended to use the WebSocket User channel to retrieve all OPEN orders.
//
//  2. start_date and end_date parameters don’t apply to open orders.
//
//     You cannot query for OPEN orders with other order types.
//     # Allowed
//     /orders/historical/batch?order_status=OPEN
//     /orders/historical/batch?order_status=CANCELLED,EXPIRED
//
//     # Not allowed
//     /orders/historical/batch?order_status=OPEN,CANCELLED
func (s *OrdersService) List(ctx context.Context, options *ListOrdersOptions) (*ListOrdersResponse, error) {
	url := s.client.baseURL + "/api/v3/brokerage/orders/historical/batch"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate list orders HTTP request: %w", err)
	}

	if options != nil {
		qs, err := query.Values(options)
		if err != nil {
			return nil, fmt.Errorf("failed to convert ListOrdersOptions to query string: %w", err)
		}

		req.URL.RawQuery = qs.Encode()
	}

	var orderResp ListOrdersResponse
	err = s.client.do(req, http.StatusOK, &orderResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orders: %w", err)
	}

	return &orderResp, err
}
