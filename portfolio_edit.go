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

type EditPortfolioOptions struct {
	PortfolioUUID string `json:"-"`    // The portfolio UUID.
	Name          string `json:"name"` // The name of the portfolio.
}

// Edit a portfolio.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_editportfolio
func (s *PortfoliosService) Edit(
	ctx context.Context,
	opts *EditPortfolioOptions,
) (*Portfolio, error) {

	u := fmt.Sprintf(
		"%s/api/v3/brokerage/portfolios/%s",
		s.client.baseURL,
		id.String(),
	)

	b, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to marshal edit portfolio options to JSON: %w",
			err,
		)
	}

	var portfolio Portfolio

	err = s.client.put(ctx, u, bytes.NewReader(b), &portfolio)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to edit protfolio '%s': %w",
			id.String(),
			err,
		)
	}

	return &portfolio, nil
}
