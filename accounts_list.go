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
	PaginatedResponse
}

type ListAccountsOptions struct {
	PaginationOptions // Options to control API response pagination.
	// Deprecated.
	// Only returns the accounts matching the portfolio ID. Only applicable
	// for legacy keys. CDP keys will default to the key's permissioned
	// portfolio.
	RetailPortfolioID *string `url:"retail_portfolio_id,omitempty"`
}

// Get a list of authenticated Advanced Trade accounts for the current user.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getaccounts
func (s *AccountsService) List(
	ctx context.Context,
	options *ListAccountsOptions,
) (*ListAccountsResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/accounts"

	var accountsResp ListAccountsResponse
	err := s.client.get(ctx, u, options, &accountsResp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch list of authenticated accounts for the current "+
				"user: %w",
			err,
		)
	}

	return &accountsResp, err
}
