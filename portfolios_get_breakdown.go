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

	"github.com/google/uuid"
)

type GetPortfolioBreakdownOptions struct {
	PortfolioUUID uuid.UUID `url:"-"`                  // The portfolio UUID.
	Currency      *string   `url:"currency,omitempty"` // The currency symbol (e.g. USD).
}

type portfolioBreakdownResponse struct {
	Breakdown PortfolioBreakdown `json:"breakdown"` // PortfolioBreakdown is a breakdown of a portfolio, all balances, and all positions within the portfolio.
}

// Get the breakdown of a portfolio.
//   - All balances
//   - All positions
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getportfoliobreakdown
func (s *PortfoliosService) GetPortfolioBreakdown(
	ctx context.Context,
	opts *GetPortfolioBreakdownOptions,
) (*PortfolioBreakdown, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/portfolios/%s",
		s.client.baseURL,
		opts.PortfolioUUID,
	)

	var resp portfolioBreakdownResponse

	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get portfolio breakdown for '%s': %w",
			opts.PortfolioUUID,
			err,
		)
	}

	return &resp.Breakdown, err
}
