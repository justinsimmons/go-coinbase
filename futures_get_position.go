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

type getFuturesPosition struct {
	Position *FuturesPosition `json:"position"`
}

// Get positions for a specific CFM product.
// The product id will be a ticker symbol (e.g. 'BIT-28JUL23-CDE').
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_getfcmposition
func (s *FuturesService) GetPosition(
	ctx context.Context,
	productID string,
) (*FuturesPosition, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/cfm/positions/%s",
		s.client.baseURL,
		productID,
	)

	var resp getFuturesPosition

	err := s.client.get(ctx, u, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get futures position: %w", err)
	}

	return resp.Position, nil
}
