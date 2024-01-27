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

type createPortfolioRequest struct {
	Name string `json:"name"`
}

type createPortfolioResponse struct {
	Portfolio *Portfolio `json:"portfolio"`
}

// Create creates a portfolio.
func (s *PortfoliosService) Create(ctx context.Context, name string) (*Portfolio, error) {
	b, err := json.Marshal(createPortfolioRequest{Name: name})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create portfolio request body to JSON: %w", err)
	}

	var portfolioResp createPortfolioResponse

	err = s.client.post(ctx, s.client.baseURL+"/api/v3/brokerage/portfolios", bytes.NewBuffer(b), &portfolioResp)
	if err != nil {
		return nil, fmt.Errorf("failed to create portfolio: %w", err)
	}

	return portfolioResp.Portfolio, nil
}
