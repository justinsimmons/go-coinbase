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

type ToggleMultiAssetCollateralOptions struct {
	PortfolioUUID               string `json:"portfolio_uuid,omitempty"`                 // The portfolio UUID.
	MultiAssetCollateralEnabled bool   `json:"multi_asset_collateral_enabled,omitempty"` // Enable or disable Multi Asset Collateral.
}

type ToggleMultiAssetCollateralResponse struct {
	MultiAssetCollateralEnabled bool `json:"multi_asset_collateral_enabled,omitempty"` // Enable or disable Multi Asset Collateral.
}

// Enable or Disable Multi Asset Collateral for a given Portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_intxmultiassetcollateral
func (s *PerpetualsService) ToggleMultiAssetCollateral(
	ctx context.Context,
	options ToggleMultiAssetCollateralOptions,
) (*ToggleMultiAssetCollateralResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/intx/multi_asset_collateral"

	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal ToggleMultiAssetCollateralOptions to JSON: %w",
			err,
		)
	}

	var resp ToggleMultiAssetCollateralResponse
	err = s.client.post(ctx, u, bytes.NewBuffer(b), &resp)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
	}

	return &resp, err
}
