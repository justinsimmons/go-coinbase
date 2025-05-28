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

// Expected response from the Get Perpetuals Portfolio Summary API.
type GetPerpetualsPortfolioSummaryResponse struct {
	Portfolios []PerpetualsPortfolio `json:"portfolios,omitempty"`
	Summary    *PortfolioSummary     `json:"portfolio_summary,omitempty"`
}

// Get a summary of your Perpetuals portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getintxportfoliosummary
func (s *PerpetualsService) GetPortfolioSummary(
	ctx context.Context,
	portfolioID string,
) (*GetPerpetualsPortfolioSummaryResponse, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/intx/%s",
		s.client.baseURL,
		portfolioID,
	)

	var resp GetPerpetualsPortfolioSummaryResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch perpetuals portfolio summary for '%s': %w",
			portfolioID,
			err,
		)
	}

	return &resp, nil
}
