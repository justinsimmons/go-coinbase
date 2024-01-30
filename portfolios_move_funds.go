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

	"github.com/google/uuid"
)

type PortfolioMoveFundsOptions struct {
	Funds               Funds     `json:"funds"`                 // Represents a monetary amount.
	SourcePortfolioUUID uuid.UUID `json:"source_portfolio_uuid"` // UUID of the portfolio to transfer funds from.
	TargetPortfolioUUID uuid.UUID `json:"target_portfolio_uuid"` // UUID of the portfolio to transfer funds to.
}

type PorfoliosMoveFundsResponse struct {
	SourcePortfolioUUID *uuid.UUID `json:"source_portfolio_uuid"` // UUID of the portfolio the funds were transfered from.
	TargetPortfolioUUID *uuid.UUID `json:"target_portfolio_uuid"` // UUID of the portfolio the funds were transfered to.
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
