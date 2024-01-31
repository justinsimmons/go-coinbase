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
	"time"
)

type SweepStatus string

const (
	SweepStatusUnknown    SweepStatus = "UNKNOWN_FCM_SWEEP_STATUS"
	SweepStatusPending    SweepStatus = "PENDING"
	SweepStatusProcessing SweepStatus = "PROCESSING"
)

type FuturesSweep struct {
	ID              *string `json:"id"` // The ID of the sweep request scheduled.
	RequestedAmount struct {
		Value    *string `json:"value"`
		Currency *string `json:"currency"`
	} `json:"requested_amount"` //The requested sweep amount.
	ShouldSweepAll *bool        `json:"should_sweep_all"` // True if the request was to sweep all available funds from your CFM futures account
	Status         *SweepStatus `json:"status"`           // A pending sweep is a sweep that has not started processing and can be cancelled. A processing sweep is a sweep that is currently being processed and cannot be cancelled.
	ScheduledTime  *time.Time   `json:"scheduled_time"`   // The timestamp at which the sweep request was submitted.
}

type listFuturesSweepsResponse struct {
	Sweeps []FuturesSweep `json:"sweeps"`
}

// ListSweeps gets information on your pending and/or processing requests to sweep funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet.
//   - A pending sweep is a sweep that has not started processing and can be cancelled.
//   - A processing sweep is a sweep that is currently being processed and cannot be cancelled.
//
// Once a sweep is complete, it longer appears in the list of sweeps
func (s *FuturesService) ListSweeps(ctx context.Context) ([]FuturesSweep, error) {
	var sweepsResp listFuturesSweepsResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/cfm/sweeps", nil, &sweepsResp)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of sweeps: %w", err)
	}

	return sweepsResp.Sweeps, nil
}
