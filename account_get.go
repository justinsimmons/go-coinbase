package coinbase

import (
	"context"
	"fmt"
	"net/http"
)

type getAccountResponse struct {
	Account Account `json:"account"`
}

// Get retrieves a list of information about an account, given an account UUID.
func (s *AccountService) Get(ctx context.Context, id string) (*Account, error) {
	url := fmt.Sprintf("%s/api/v3/brokerage/accounts/%s", s.client.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get account HTTP request: %w", err)
	}

	var accountResp getAccountResponse
	err = s.client.do(req, http.StatusOK, &accountResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account '%s' for the current user: %w", id, err)
	}

	return &accountResp.Account, err
}
