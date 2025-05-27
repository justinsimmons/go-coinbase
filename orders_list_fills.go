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

type ListOrderFillsOptions struct {
	OrderIDs          []string   `url:"order_ids,omitempty"`                // The ID(s) of order(s).
	TradeIDs          []string   `url:"trade_ids,omitempty"`                // The ID(s) of the trades of fills.
	ProductIDs        []string   `url:"product_ids,omitempty"`              // The ID(s) of the product(s) to filter fills by.
	StartSequenceTime *time.Time `url:"start_sequence_timestamp,omitempty"` // Only fills with a trade time after the specified start date are returned.
	EndSequenceTime   *time.Time `url:"end_sequence_timestamp,omitempty"`   // Only fills with a trade time before the specified end date are returned.
	RetailPortfolioID *string    `url:"retail_portfolio_id,omitempty"`      // (Deprecated) Only orders matching this retail portfolio id are returned. Only applicable for legacy keys. CDP keys will default to the key's permissioned portfolio.
	Limit             *int32     `url:"limit,omitempty"`                    // The number of orders to display per page (no default amount). If has_next is true, additional pages of orders are available to be fetched. Use the cursor parameter to start on a specified page.
	Cursor            *string    `url:"cursor,omitempty"`                   // For paginated responses, returns all responses that come after this value.
	SortBy            *string    `url:"sort_by,omitempty"`                  // Sort results by a field, results use unstable pagination. Default is to sort by creation time.
}

type ListFillsResponse struct {
	Fills  []Fill  `json:"fills"`  // All fills matching the filters.
	Cursor *string `json:"cursor"` // Cursor for paginating. Users can use this string to pass in the next call to this endpoint, and repeat this process to fetch all fills through pagination.
}

// Get a list of fills filtered by optional query parameters (product_id, order_id, etc).
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getfills
func (s *OrdersService) ListFills(
	ctx context.Context,
	options *ListOrderFillsOptions,
) (*ListFillsResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/historical/fills"

	var resp ListFillsResponse

	err := s.client.get(ctx, u, &options, &resp)
	if err != nil {
		err = fmt.Errorf("failed to list historical fills: %w", err)
	}

	return &resp, err
}
