package coinbase

import (
	"context"
	"fmt"
)

type listPaymentMethodsResponse struct {
	PaymentMethods []PaymentMethod `json:"payment_methods"`
}

// Get a list of payment methods for the current user.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpaymentmethods/
func (s *PaymentMethodsService) List(ctx context.Context) ([]PaymentMethod, error) {
	u := s.client.baseURL + "/api/v3/brokerage/payment_methods"

	var resp listPaymentMethodsResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all payment methods for the current user: %w", err)
	}

	return resp.PaymentMethods, nil

}
