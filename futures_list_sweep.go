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

type listFuturesSweepsResponse struct {
	Sweeps []FuturesSweep `json:"sweeps"`
}

// Get pending and processing sweeps of funds from FCM wallet to USD Spot wallet.
//   - A pending sweep is a sweep that has not started processing and can be cancelled.
//   - A processing sweep is a sweep that is currently being processed and cannot be cancelled.
//
// Once a sweep is complete, it longer appears in the list of sweeps.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getfcmsweeps
func (s *FuturesService) ListSweeps(
	ctx context.Context,
) ([]FuturesSweep, error) {

	u := s.client.baseURL + "/api/v3/brokerage/cfm/sweeps"

	var resp listFuturesSweepsResponse

	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of sweeps: %w", err)
	}

	return resp.Sweeps, nil
}
