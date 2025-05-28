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

// Expected response from the Get Perpetuals Portfolio Balances API.
type getPerpetualsPortfolioBalancesResponse struct {
	PortfolioBalances []PerpetualsPortfolioBalance `json:"portfolio_balances,omitempty"`
}

// Get a list of asset balances on Intx for a given Portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getintxbalances
func (s *PerpetualsService) GetPortfolioBalances(
	ctx context.Context,
	porfilioID string,
) ([]PerpetualsPortfolioBalance, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/intx/balances/%s",
		s.client.baseURL,
		porfilioID,
	)

	var resp getPerpetualsPortfolioBalancesResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch perpetuals portfolio balances for '%s': %w",
			porfilioID,
			err,
		)
	}

	return resp.PortfolioBalances, nil
}
