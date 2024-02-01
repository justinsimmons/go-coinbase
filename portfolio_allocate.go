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

type AllocatePortfolioOptions struct {
	PortfolioUUID uuid.UUID `json:"portfolio_uuid"` // The unique identifier for the perpetuals portfolio.
	Symbol        string    `json:"symbol"`         // The product_id for which funds must be allocated.
	Amount        string    `json:"amount"`         // The value of funds to be allocated for a specific isolated position.
	Currency      string    `json:"currency"`       // The currency of funds to be allocated for a specific isolated position.
}

// Allocate allocates more funds to an isolated position in your Perpetuals portfolio.
func (s *PortfoliosService) Allocate(ctx context.Context, options AllocatePortfolioOptions) error {
	b, err := json.Marshal(options)
	if err != nil {
		return fmt.Errorf("failed to marshal AllocatePortfolioOptions to JSON: %w", err)
	}

	// Response is an empty object. Scan response to map and discard.
	var resp map[string]any

	err = s.client.post(ctx, s.client.baseURL+"/api/v3/brokerage/intx/allocate", bytes.NewReader(b), &resp)
	if err != nil {
		return fmt.Errorf("failed to allocate portfolio: %w", err)
	}

	return nil
}
