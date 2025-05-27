// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type PreviewOrderOptions struct {
	ProductID                  string              `json:"product_id"`                             // The trading pair (e.g. 'BTC-USD').
	Side                       Side                `json:"side"`                                   // The side of the market that the order is on (e.g. 'BUY', 'SELL').
	OrderConfiguration         OrderConfiguration  `json:"order_configuration"`                    // The configuration of the order (e.g. the order type, size, etc).
	Leverage                   *string             `json:"leverage,omitempty"`                     // The amount of leverage for the order (default is 1.0).
	MarginType                 *MarginType         `json:"margin_type,omitempty"`                  // Margin Type for this order (default is CROSS).
	RetailPortfolioID          *string             `json:"retail_portfolio_id,omitempty"`          // (Deprecated) The ID of the portfolio to associate the order with. Only applicable for legacy keys. CDP keys will default to the key's permissioned portfolio.
	AttachedOrderConfiguration *OrderConfiguration `json:"attached_order_configuration,omitempty"` // The configuration of the attached order. Only TriggerBracketGtc is eligible. Size field must be omitted as the size of the attached order is the same as that of the parent order.
}

type TwapBucketMetadata struct {
	NumberBuckets  *string        `json:"number_buckets,omitempty"`  // The number of smaller buckets/suborders over which the entire order will be broken into. Each suborder will be executed over a duration calculated based on the end_time.
	BucketSize     *string        `json:"bucket_size,omitempty"`     // The size of each suborder. bucket_size multiplied by number_buckets should match the size of the entire twap order).
	BucketDuration *time.Duration `json:"bucket_duration,omitempty"` // The duration over which each sub order was executed.
}

type PreviewOrderResponse struct {
	OrderTotal                 string             `json:"order_total"`
	ComissionTotal             string             `json:"comission_total"` // Currency amount of the applied commission (so not the rate that was used on input).
	Errors                     []string           `json:"errs"`            // List of potential failure reasons were this order to be submitted.
	Warnings                   []string           `json:"warning"`
	QuoteSize                  string             `json:"quote_size"` // The amount of the second Asset in the Trading Pair. For example, on the BTC/USD Order Book, USD is the Quote Asset.
	BaseSize                   string             `json:"base_size"`  // The amount of the first Asset in the Trading Pair. For example, on the BTC-USD Order Book, BTC is the Base Asset.
	BestBid                    string             `json:"best_bid"`
	BestAsk                    string             `json:"best_ask"`
	IsMax                      bool               `json:"is_max"` // Indicates whether tradable_balance should be set to the maximum amount.
	OrderMarginTotal           *string            `json:"order_margin_total,omitempty"`
	Leverage                   *string            `json:"leverage,omitempty"` // The amount of leverage for the order (default is 1.0).
	LongLeverage               *string            `json:"long_leverage,omitempty"`
	ShortLeverage              *string            `json:"short_leverage,omitempty"`
	Slippage                   *string            `json:"slippage,omitempty"`
	PreviewID                  *string            `json:"preview_id,omitempty"`
	CurrentLiquidationBuffer   *string            `json:"current_liquidation_buffer,omitempty"`
	ProjectedLiquidationBuffer *string            `json:"projected_liquidation_buffer,omitempty"`
	MaxLeverage                *string            `json:"max_leverage,omitempty"`
	PnlConfiguration           PnlConfiguration   `json:"pnl_configuration,omitempty"` // Expected PNL of an order. This value is an estimate and does not take into account fees and slippage.
	TwapBucketMetadata         TwapBucketMetadata `json:"twap_bucket_metadata,omitempty"`
	PositionNotionalLimit      *string            `json:"position_notional_limit,omitempty"`
}

// Preview an order.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_previeworder
func (s *OrdersService) Preview(
	ctx context.Context,
	options CreateOrderOptions,
) (*PreviewOrderResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/preview"

	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal PreviewOrderRequest to JSON: %w",
			err,
		)
	}

	var resp PreviewOrderResponse
	err = s.client.post(ctx, u, bytes.NewBuffer(b), &resp)
	if err != nil {
		err = fmt.Errorf("failed to preview order: %w", err)
	}

	return &resp, err
}
