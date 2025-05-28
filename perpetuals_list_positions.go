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

// Expected response from the List Perpetuals positions API.
type ListPerpetualsPositionsResponse struct {
	Positions []PerpetualsPosition `json:"positions,omitempty"`
	// Alright Coinbase we need to have a chat about your API's...
	// Why can't you be consistent in the SAME domain!??? A perpetuals summary
	// should be the same across all the API's.
	Summary struct {
		AggregatedPNL *AvailableBalance `json:"aggregated_pnl,omitempty"`
	} `json:"summary,omitempty"`
}

// Get a list of open positions in your Perpetuals portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getintxpositions
func (s *PerpetualsService) ListPositions(
	ctx context.Context,
	portfolioUUID string,
) (*ListPerpetualsPositionsResponse, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/intx/positions/%s",
		s.client.baseURL,
		portfolioUUID,
	)

	var resp ListPerpetualsPositionsResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch list of perpetuals portfolio positions for "+
				"'%s': %w",
			portfolioUUID,
			err,
		)
	}

	return &resp, nil
}
