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
