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

type GetCurrentMarginWindowOptions struct {
	MarginProfileType *MarginProfileType `url:"margin_profile_type,omitempty"` // The margin profile type for your account.
}

type GetCurrentMarginWindowResponse struct {
	MarginWindow                                *MarginWindow `json:"margin_window,omitempty"`                                    // Margin window.
	IsIntradayMarginKillswitchEnabled           bool          `json:"is_intraday_margin_killswitch_enabled,omitempty"`            // True if intraday margin killswitch is enabled.
	IsIntradayMarginEnrollmentKillswitchEnabled bool          `json:"is_intraday_margin_enrollment_killswitch_enabled,omitempty"` // True if intraday margin enrollment killswitch is enabled.
}

// Get the futures current margin window.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getcurrentmarginwindow
func (s *FuturesService) GetCurrentMarginWindow(
	ctx context.Context,
	opts *GetCurrentMarginWindowOptions,
) (*GetCurrentMarginWindowResponse, error) {

	u := s.client.baseURL +
		"/api/v3/brokerage/cfm/intraday/current_margin_window"

	var resp GetCurrentMarginWindowResponse

	err := s.client.get(ctx, u, opts, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to get current margin window: %w",
			err,
		)
	}

	return &resp, nil
}
