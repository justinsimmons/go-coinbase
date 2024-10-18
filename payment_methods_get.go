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

// Expected response from the Get Payment Method API.
type getPaymentMethodResponse struct {
	PaymentMethod PaymentMethod `json:"payment_method"`
}

// Get information about a payment method for the current user.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpaymentmethod/
func (s *PaymentMethodsService) Get(ctx context.Context, id string) (*PaymentMethod, error) {
	u := fmt.Sprintf("%s/api/v3/brokerage/payment_methods/%s", s.client.baseURL, id)

	var resp getPaymentMethodResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch payment method '%s' for the current user: %w", id, err)
	}

	return &resp.PaymentMethod, nil
}
