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

type ListAccountsResponse struct {
	Accounts []Account `json:"accounts"`
	// Whether there are additional pages for this query.
	HasNext bool `json:"has_next"`
	// Cursor for paginating. Users can use this string to pass in the next call to this
	// endpoint, and repeat this process to fetch all accounts through pagination.
	Cursor *string `json:"cursor"`
	// Number of accounts returned
	Size *int32 `json:"size"`
}

type AccountListOptions struct {
	// A pagination limit with default of 49 and maximum of 250.
	// If has_next is true, additional orders are available to be fetched with pagination and the cursor value
	// in the response can be passed as cursor parameter in the subsequent request.
	Limit *int32 `url:"limit,omitempty"`
	// Cursor used for pagination. When provided, the response returns responses after this cursor.
	Cursor *string `url:"cursor,omitempty"`
}

// List retrieves a list of authenticated accounts for the current user.
// https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccounts
func (s *AccountService) List(ctx context.Context, options *AccountListOptions) (*ListAccountsResponse, error) {
	var accountsResp ListAccountsResponse
	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/accounts", options, &accountsResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of authenticated accounts for the current user: %w", err)
	}

	return &accountsResp, err
}
