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

type ScheduleSweepOptions struct {
	USDAmmount *string `json:"usd_amount,omitempty"` // The amount of USD to be swept. By default, sweeps all available excess funds.
}

type ScheduleSweepResponse struct {
	Success bool `json:"success"`
}

// Schedules a sweep of funds from FCM wallet to USD Spot wallet.
//   - Sweep requests submitted before 5PM ET each day are processed the following business day.
//   - Sweep requests submitted after 5PM ET each day are processed in 2 business days.
//
// You can have at most one pending sweep request at a time.
//
// Market movements related to your open positions may impact the final amount
// that is transferred into your spot account. The final funds transferred,
// up to your specified amount, depend on the available excess in your futures
// account.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_schedulefcmsweep
func (s *FuturesService) ScheduleSweep(
	ctx context.Context,
	opts ScheduleSweepOptions,
) (*ScheduleSweepResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/cfm/sweeps/schedule"

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal ScheduleSweepOptions to JSON: %w",
			err,
		)
	}

	var resp ScheduleSweepResponse

	err = s.client.post(ctx, u, bytes.NewReader(b), &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule sweep: %w", err)
	}

	return &resp, nil
}
