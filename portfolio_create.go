// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type CreatePortfolioOptions struct {
	Name string `json:"name"`
}

type createPortfolioResponse struct {
	Portfolio *Portfolio `json:"portfolio"`
}

// Create a portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_createportfolio
func (s *PortfoliosService) Create(
	ctx context.Context,
	opts *CreatePortfolioOptions,
) (*Portfolio, error) {

	u := s.client.baseURL + "/api/v3/brokerage/portfolios"

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal create portfolio request body to JSON: %w",
			err,
		)
	}

	var portfolioResp createPortfolioResponse

	err = s.client.post(ctx, u, bytes.NewBuffer(b), &portfolioResp)
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio: %w", err)
	}

	return portfolioResp.Portfolio, nil
}
