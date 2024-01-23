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
	PortfolioType *PortfolioType `url:"portfolio_type,omitempty"`
}

type ListPortfoliosResponse struct {
	Portfolios []Portfolio `json:"portfolios"`
}

// List gets a list of all portfolios of a user.
func (s *PortfoliosService) List(ctx context.Context, options *ListPortfoliosOptions) (*ListPortfoliosResponse, error) {
	var portfolios ListPortfoliosResponse
	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/portfolios", options, &portfolios)
	if err != nil {
		err = fmt.Errorf("failed to get list of portfolios: %w", err)
	}

	return &portfolios, err
}
