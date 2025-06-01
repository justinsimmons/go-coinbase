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

type ListPortfoliosOptions struct {
	PortfolioType *PortfolioType `url:"portfolio_type,omitempty"` // Only returns portfolios matching this portfolio type.
}

type ListPortfoliosResponse struct {
	Portfolios []Portfolio `json:"portfolios"`
}

// Get all portfolios of a user.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getportfolios
func (s *PortfoliosService) List(
	ctx context.Context,
	options *ListPortfoliosOptions,
) ([]Portfolio, error) {

	u := s.client.baseURL + "/api/v3/brokerage/portfolios"

	var resp ListPortfoliosResponse

	err := s.client.get(ctx, u, options, &resp)
	if err != nil {
		err = fmt.Errorf("failed to get list of portfolios: %w", err)
	}

	return resp.Portfolios, err
}
