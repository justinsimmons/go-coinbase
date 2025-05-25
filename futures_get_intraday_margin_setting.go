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

type getIntradayMarginSettingResponse struct {
	Setting *IntradayMarginSetting `json:"setting,omitempty"`
}

// Get the futures intraday margin setting.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getintradaymarginsetting
func (s *FuturesService) GetIntradayMarginSetting(
	ctx context.Context,
) (IntradayMarginSetting, error) {

	u := s.client.baseURL + "/api/v3/brokerage/cfm/intraday/margin_setting"

	var resp getIntradayMarginSettingResponse

	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return IntradayMarginSettingUnspecified, fmt.Errorf(
			"failed to get intraday margin setting: %w",
			err,
		)
	}

	// Should never happen but prevents a null pointer deref on an API error.
	if resp.Setting == nil {
		return IntradayMarginSettingUnspecified, fmt.Errorf(
			"coinbase API response successful, but missing intraday margin " +
				"setting",
		)
	}

	return *resp.Setting, nil
}
