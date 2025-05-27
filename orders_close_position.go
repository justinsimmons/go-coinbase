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

type ClosePositionOptions struct {
	ClientOrderID string  `json:"client_order_id"` // The unique ID provided for the order (used for identification purposes).
	ProductID     string  `json:"product_id"`      // The trading pair (e.g. 'BIT-28JUL23-CDE').
	Size          *string `json:"size,omitempty"`  // The amount of contracts that should be closed.
}

type ClosePositionResponse struct {
	Success            bool                 `json:"success"` // Whether the order was created.
	SuccessResponse    OrderSuccessMetadata `json:"success_response"`
	ErrorResponse      OrderErrorMetadata   `json:"error_response"`
	OrderConfiguration *OrderConfiguration  `json:"order_configuration"` // The configuration of the order (e.g. the order type, size, etc).
}

// Places an order to close any open positions for a specified product_id.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_closeposition
func (s *OrdersService) ClosePosition(
	ctx context.Context,
	options *ClosePositionOptions,
) (*ClosePositionResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/close_position"

	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal ClosePositionOptions to JSON: %w",
			err,
		)
	}

	var resp ClosePositionResponse
	err = s.client.post(ctx, u, bytes.NewBuffer(b), &resp)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
	}

	return &resp, err
}
