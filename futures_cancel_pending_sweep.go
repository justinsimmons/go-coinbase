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
	"net/http"
)

type cancelPendingFuturesSweepResponse struct {
	Success bool `json:"success"`
}

// CancelPendingSweep cancels your pending sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet.
// Returns true if all sweeps are successfully canceled.
func (s *FuturesService) CancelPendingSweep(ctx context.Context) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, s.client.baseURL+"/api/v3/brokerage/cfm/sweeps", nil)
	if err != nil {
		return false, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var resp cancelPendingFuturesSweepResponse
	err = s.client.do(req, http.StatusOK, &resp)
	if err != nil {
		return false, fmt.Errorf("failed to cancel pending futures sweep: %w", err)
	}

	return resp.Success, nil
}
