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

type TradeType string

const (
	TradeTypeFill       TradeType = "FILL"
	TradeTypeReversal   TradeType = "REVERSAL"
	TradeTypeCorrection TradeType = "CORRECTION"
	TradeTypeSynthetic  TradeType = "SYNTHETIC"
)

type LiquidityIndicator string

const (
	LiquidityIndicatorUnknown LiquidityIndicator = "UNKNOWN_LIQUIDITY_INDICATOR"
	LiquidityIndicatorMaker   LiquidityIndicator = "MAKER"
	LiquidityIndicatorTaker   LiquidityIndicator = "TAKER"
)

type ListOrderFillsOptions struct {
	OrderID           *string    `url:"order_id,omitempty"`                 // ID of order.
	ProductID         *string    `url:"product_id,omitempty"`               // The ID of the product this order was created for.
	StartSequenceTime *time.Time `url:"start_sequence_timestamp,omitempty"` // Start date. Only fills with a trade time at or after this start date are returned.
	EndSequenceTime   *time.Time `url:"end_sequence_timestamp,omitempty"`   // End date. Only fills with a trade time before this start date are returned.
	Limit             *int64     `url:"limit,omitempty"`                    // Maximum number of fills to return in response. Defaults to 100.
	Cursor            *string    `url:"cursor,omitempty"`                   // Cursor used for pagination. When provided, the response returns responses after this cursor.
}

type Fill struct {
	EntryID            *string             `json:"entry_id"`            // Unique identifier for the fill.
	TradeID            *string             `json:"trade_id"`            // ID of the fill -- unique for all `FILL` trade_types but not unique for adjusted fills.
	OrderID            *string             `json:"order_id"`            // ID of the order the fill belongs to.
	TradeTime          *time.Time          `json:"trade_time"`          // Time at which this fill was completed.
	TradeType          *TradeType          `json:"trade_type"`          // String denoting what type of fill this is. Regular fills have the value `FILL`. Adjusted fills have possible values `REVERSAL`, `CORRECTION`, `SYNTHETIC`.
	Price              *string             `json:"price"`               // Price the fill was posted at.
	Size               *string             `json:"size"`                // Amount of order that was transacted at this fill.
	Commission         *string             `json:"commission"`          //  Fee amount for fill.
	ProductID          *string             `json:"product_id"`          // The product this order was created for.
	SequenceTimestamp  *time.Time          `json:"sequence_timestamp"`  // Time at which this fill was posted.
	LiquidityIndicator *LiquidityIndicator `json:"liquidity_indicator"` // Possible values: [UNKNOWN_LIQUIDITY_INDICATOR, MAKER, TAKER]
	SizeInQuote        *bool               `json:"size_in_quote"`       // Whether the order was placed with quote currency.
	UserID             *string             `json:"user_id"`             // User that placed the order the fill belongs to.
	Side               *Side               `json:""`                    // Side the fill is on [BUY, SELL].
}

type ListFillsResponse struct {
	Fills  []Fill  `json:"fills"`  // All fills matching the filters.
	Cursor *string `json:"cursor"` // Cursor for paginating. Users can use this string to pass in the next call to this endpoint, and repeat this process to fetch all fills through pagination.
}

// Get a list of fills filtered by optional query parameters (product_id, order_id, etc).
func (s *OrdersService) ListFills(ctx context.Context, options *ListOrderFillsOptions) (*ListFillsResponse, error) {
	var fills ListFillsResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/orders/historical/fills", &options, &fills)
	if err != nil {
		err = fmt.Errorf("failed to list historical fills: %w", err)
	}

	return &fills, err
}
