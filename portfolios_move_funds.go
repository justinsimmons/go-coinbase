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

type Funds struct {
	Value    string `json:"value"`    // These two values fully represent the monetary amount. Non-localized amount in decimal notation (e.g. "1.234").
	Currency string `json:"currency"` // Currency symbol (USD, BTC, etc). Not an asset UUID.
}

type PortfolioMoveFundsOptions struct {
	Funds             Funds  `json:"funds"`                 // Represents a monetary amount.
	SourcePortfolioID string `json:"source_portfolio_uuid"` // UUID of the portfolio to transfer funds from.
	TargetPortfolioID string `json:"target_portfolio_uuid"` // UUID of the portfolio to transfer funds to.
}

type PorfoliosMoveFundsResponse struct {
	SourcePortfolioID *string `json:"source_portfolio_uuid"` // UUID of the portfolio the funds were transfered from.
	TargetPortfolioID *string `json:"target_portfolio_uuid"` // UUID of the portfolio the funds were transfered to.
}

// MoveFunds transfers funds between portfolios.
func (s *PortfoliosService) MoveFunds(ctx context.Context, options PortfolioMoveFundsOptions) (*PorfoliosMoveFundsResponse, error) {
	b, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal move portfolio funds request body to JSON: %w", err)
	}

	var resp PorfoliosMoveFundsResponse

	err = s.client.post(ctx, s.client.baseURL+"/api/v3/brokerage/portfolios/move_funds", bytes.NewReader(b), &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to move portfolio funds: %w", err)
	}

	return &resp, nil
}
