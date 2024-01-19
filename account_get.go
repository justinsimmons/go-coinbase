package cb

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

	req, err := s.client.newRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate get account request: %w", err)
	}

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch account '%s' for the current user: %w", id, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, handleRequestError(resp)
	}

	var accountResp getAccountResponse
	err = unmarshal(resp, &accountResp)

	return &accountResp.Account, err
}
