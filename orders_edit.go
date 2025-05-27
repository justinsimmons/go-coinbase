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

type EditOrderOptions struct {
	OrderID string `json:"order_id"` // ID of order to edit.
	Price   string `json:"price"`    // New price for order.
	Size    string `json:"size"`     // New size for order.
}

type EditOrderResponse struct {
	Success bool             `json:"success"` // Whether the order edit request was placed.
	Errors  []EditOrderError `json:"errors"`  // Details of any errors that may have occurred.
}

func (s *OrdersService) edit(
	ctx context.Context,
	url string,
	options *EditOrderOptions,
) (*EditOrderResponse, error) {

	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal edit order options to JSON: %w",
			err,
		)
	}

	var resp EditOrderResponse
	err = s.client.post(ctx, url, bytes.NewBuffer(b), &resp)
	if err != nil {
		err = fmt.Errorf("failed to edit order: %w", err)
	}

	return &resp, err
}

// Preview an edit order request with a specified new size, or new price.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_previeweditorder
func (s *OrdersService) EditPreview(
	ctx context.Context,
	options *EditOrderOptions,
) (*EditOrderResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/edit_preview"

	return s.edit(ctx, u, options)
}

// Edit an order with a specified new size, or new price.
//
// Order Priority:
//   - A client can only send an Edit Order request after the previous request for the same order has been fully processed.
//   - CAUTION: You lose your place in line if you increase size or increase/decrease price.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_editorder
func (s *OrdersService) Edit(
	ctx context.Context,
	options *EditOrderOptions,
) (*EditOrderResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/orders/edit"

	return s.edit(ctx, u, options)
}
