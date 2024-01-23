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
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetTransactionsSummaryOptions struct {
	ProductType        *string             `url:"product_type,omitempty"`
	ContractExpiryType *ContractExpiryType `url:"contract_expiry_type,omitempty"`
}

type FeeTier struct {
	PricingTier  *string `json:"pricing_tier"`   // Pricing tier for user, determined by notional (USD) volume.
	USDFrom      *string `json:"usd_from"`       // Lower bound (inclusive) of pricing tier in notional volume.
	USDTo        *string `json:"usd_to"`         // Upper bound (exclusive) of pricing tier in notional volume.
	TakerFeeRate *string `json:"taker_fee_rate"` // Taker fee rate, applied if the order takes liquidity.
	MakerFeeRate *string `json:"maker_fee_rate"` // Maker fee rate, applied if the order creates liquidity.
	AOPFrom      *string `json:"aop_from"`       // Lower bound (inclusive) of pricing tier in usd of total assets on platform.
	AOPTO        *string `json:"aop_to"`         // Upper bound (exclusive) of pricing tier in usd of total assets on platform.
}

type GetTransactionsSummaryResponse struct {
	TotalVolume         float64 `json:"total_volume"` // Total volume across assets, denoted in USD.
	TotalFees           float64 `json:"total_fees"`   // Total fees across assets, denoted in USD.
	FeeTier             FeeTier `json:"fee_tier"`
	GoodsAndServicesTax struct {
		Rate *string `json:"rate"`
		Type *string `json:"type"` // Possible values: [INCLUSIVE, EXCLUSIVE] // TODO: enum.
	} `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume *float64 `json:"advanced_trade_only_volume"` // Advanced Trade volume (non-inclusive of Pro) across assets, denoted in USD.
	AdvancedTradeOnlyFees   *float64 `json:"advanced_trade_only_fees"`   // Advanced Trade fees (non-inclusive of Pro) across assets, denoted in USD.
	CoinbaseProVolume       *float64 `json:"coinbase_pro_volume"`        // Coinbase Pro volume across assets, denoted in USD.
	CoinbaseProFees         *float64 `json:"coinbase_pro_fees"`          // Coinbase Pro fees across assets, denoted in USD.
}

// GetTransactionsSummary gets a summary of transactions with fee tiers, total volume, and fees.
func (s *FeesService) GetTransactionsSummary(ctx context.Context, options *GetTransactionsSummaryOptions) (*GetTransactionsSummaryResponse, error) {
	url := s.client.baseURL + "/api/v3/brokerage/transaction_summary"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get transactions summary HTTP request: %w", err)
	}

	if options == nil {
		qs, err := query.Values(options)
		if err != nil {
			return nil, fmt.Errorf("failed to convert GetTransactionsSummaryOptions to query string: %w", err)
		}

		req.URL.RawQuery = qs.Encode()
	}

	var summary GetTransactionsSummaryResponse
	err = s.client.do(req, http.StatusOK, &summary)
	if err != nil {
		err = fmt.Errorf("failed to fetch get transactions summary: %w", err)
	}

	return &summary, err
}
