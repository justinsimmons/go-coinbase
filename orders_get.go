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

type GetOrderOptions struct {
	OrderID            string  `url:"-"`                              // The ID of the order.
	ClientOrderID      *string `url:"client_order_id,omitempty"`      // (Deprecated) Client Order ID to fetch the order with.
	UserNativeCurrency *string `url:"user_native_currency,omitempty"` // (Deprecated) Native currency to fetch order with. Default is USD.
}

type getOrderResponse struct {
	Order Order `json:"order"`
}

// Get a single order by order ID.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_gethistoricalorder
func (s *OrdersService) Get(
	ctx context.Context,
	opts *GetOrderOptions,
) (*Order, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/orders/historical/%s",
		s.client.baseURL,
		opts.OrderID,
	)

	var resp getOrderResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch order '%s' for the current user: %w",
			opts.OrderID,
			err,
		)
	}

	return &resp.Order, err
}
