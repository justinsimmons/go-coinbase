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

type listFuturesPositionsResponse struct {
	Positions []FuturesPosition `json:"positions"`
}

// ListPositions gets a list of all open positions in CFM futures products.
func (s *FuturesService) ListPositions(ctx context.Context) ([]FuturesPosition, error) {
	var resp listFuturesPositionsResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/cfm/positions", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of futures positions: %w", err)
	}

	return resp.Positions, nil
}
