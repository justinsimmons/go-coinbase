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

type GetPerpetualsPositionOptions struct {
	PortfolioUUID string // The portfolio UUID.
	Symbol        string // The trading pair (e.g. 'BTC-PERP-INTX').
}

// Expected response from the Get Perpetuals Portfolio Summary API.
type getPerpetualsPositionResponse struct {
	Position PerpetualsPosition `json:"position,omitempty"`
}

// Get a specific open position on Intx.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getintxposition
func (s *PerpetualsService) GetPosition(
	ctx context.Context,
	opts *GetPerpetualsPositionOptions,
) (*PerpetualsPosition, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/intx/positions/%s/%s",
		s.client.baseURL,
		opts.PortfolioUUID,
		opts.Symbol,
	)

	var resp getPerpetualsPositionResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch perpetuals portfolio summary for '%s': %w",
			opts.PortfolioUUID,
			err,
		)
	}

	return &resp.Position, nil
}
