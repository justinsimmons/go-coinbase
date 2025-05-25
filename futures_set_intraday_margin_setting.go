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

type SetIntradayMarginSettingOptions struct {
	// The margin setting for the account. Describes whether the account is
	// opted in to receive increased leverage during weekdays (8am-4pm ET),
	// excluding market holidays.
	Seting *IntradayMarginSetting `json:"setting,omitempty"`
}

// Response from the set intraday margin setting API.
type SetIntradayMarginSettingResponse struct{}

// Set the futures intraday margin setting.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_setintradaymarginsetting
func (s *FuturesService) SetIntradayMarginSetting(
	ctx context.Context,
	opts *SetIntradayMarginSettingOptions,
) (*SetIntradayMarginSettingResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/cfm/intraday/margin_setting"

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal SetIntradayMarginSettingOptions to JSON: %w",
			err,
		)
	}

	var resp SetIntradayMarginSettingResponse

	err = s.client.post(ctx, u, bytes.NewReader(b), &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to set intraday margin setting: %w",
			err,
		)
	}

	return &resp, nil
}
