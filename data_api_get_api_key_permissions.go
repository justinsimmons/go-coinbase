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

	"github.com/google/uuid"
)

// Information about your CDP API key permissions.
type GetApiKeyPermissionsResponse struct {
	CanView       bool           `json:"can_view,omitempty"`       // Indicates whether the API key has view permissions.
	CanTrade      bool           `json:"can_trade,omitempty"`      // Indicates whether the API key has trade permissions.
	CanTransfer   bool           `json:"can_transfer,omitempty"`   // Indicates whether the API key has deposit/withdrawal permissions.
	PortfolioUUID *uuid.UUID     `json:"portfolio_uuid,omitempty"` // The portfolio ID associated with the API key.
	PortfolioType *PortfolioType `json:"portfolio_type,omitempty"` // The type of portfolio.
}

// Get information about your CDP API key permissions.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getapikeypermissions
func (s *DataApiService) GetApiKeyPermissions(
	ctx context.Context,
) (*GetApiKeyPermissionsResponse, error) {

	u := fmt.Sprintf("%s/api/v3/brokerage/key_permissions", s.client.baseURL)

	var resp GetApiKeyPermissionsResponse
	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to fetch api permissions for the current user: %w",
			err,
		)
	}

	return &resp, nil
}
