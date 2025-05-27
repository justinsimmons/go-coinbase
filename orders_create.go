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
)

type CreateOrderOptions struct {
	ClientOrderID              string              `json:"client_order_id"`                        // A unique ID provided for the order (used for identification purposes). If the ID provided is not unique, the order will not be created and the order corresponding with that ID will be returned instead.
	ProductID                  string              `json:"product_id"`                             // The product this order was created for e.g. 'BTC-USD'.
	Side                       *Side               `json:"side,omitempty"`                         // The side of the market that the order is on (e.g. 'BUY', 'SELL').
	OrderConfiguration         OrderConfiguration  `json:"order_configuration"`                    // Configuration of the order details.
	Leverage                   *string             `json:"leverage,omitempty"`                     // Leverage for this order; default value is "1.0".
	MarginType                 *MarginType         `json:"margin_type,omitempty"`                  // Margin Type for this order (default is CROSS).
	RetailPortfolioID          *string             `json:"retail_portfolio_id,omitempty"`          // Portfolio to place the order from and only applicable for OAuth connections. API keys will use the key's permissioned portfolio.
	PreviewID                  *string             `json:"preview_id,omitempty"`                   // Preview ID for this order, to associate this order with a preview.
	AttachedOrderConfiguration *OrderConfiguration `json:"attached_order_configuration,omitempty"` // The configuration of the attached order. Only TriggerBracketGtc is eligible. Size field must be omitted as the size of the attached order is the same as that of the parent order.
}

type CreateOrderResponse struct {
	Success            bool                 `json:"success"` // Whether the order was created.
	SuccessResponse    OrderSuccessMetadata `json:"success_response"`
	ErrorResponse      *OrderErrorMetadata  `json:"error_response,omitempty"`
	OrderConfiguration *OrderConfiguration  `json:"order_configuration,omitempty"`
}

// Create an order with a specified product_id (asset-pair), side (buy/sell), etc.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_postorder
func (s *OrdersService) Create(
	ctx context.Context,
	options CreateOrderOptions,
) (*CreateOrderResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders"

	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal OrderRequest to JSON: %w",
			err,
		)
	}

	var resp CreateOrderResponse
	err = s.client.post(ctx, u, bytes.NewBuffer(b), &resp)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
	}

	return &resp, err
}
