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
	USDAmmount *string `json:"usd_amount"` // The amount you want to sweep from your CFM futures account into your CBI spot account. Leave empty to sweep all available excess funds.
}

type ScheduleSweepResponse struct {
	Success *bool `json:"success"`
}

// ScheduleSweep schedules a sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet.
//   - Sweep requests submitted before 5PM ET each day are processed the following business day.
//   - Sweep requests submitted after 5PM ET each day are processed in 2 business days.
//
// You can have at most one pending sweep request at a time.
//
// Market movements related to your open positions may impact the final amount that is transferred into your spot account. The final funds transferred, up to your specified amount, depend on the available excess in your futures account.
func (s *FuturesService) ScheduleSweep(ctx context.Context, options ScheduleSweepOptions) (*ScheduleSweepResponse, error) {
	b, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ScheduleSweepOptions to JSON: %w", err)
	}

	var resp ScheduleSweepResponse

	err = s.client.post(ctx, s.client.baseURL+"/api/v3/brokerage/cfm/sweeps/schedule", bytes.NewReader(b), &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule sweep: %w", err)
	}

	return &resp, nil
}
