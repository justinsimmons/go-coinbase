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

type GetTransactionsSummaryOptions struct {
	ProductType        *string             `url:"product_type,omitempty"`         // Only returns the orders matching this product type. By default, returns all product types.
	ContractExpiryType *ContractExpiryType `url:"contract_expiry_type,omitempty"` // Only returns the orders matching this contract expiry type. Only applicable if product_type is set to FUTURE.
	ProductVenue       *ProductVenue       `url:"product_venue,omitempty"`        // Venue for product.
}

type GetTransactionsSummaryResponse struct {
	TotalVolume             float64              `json:"total_volume"` // Total volume across assets, denoted in USD.
	TotalFees               float64              `json:"total_fees"`   // Total fees across assets, denoted in USD.
	FeeTier                 FeeTier              `json:"fee_tier"`
	MarginRate              *MarginRate          `json:"margin_rate,omitempty"`
	GoodsAndServicesTax     *GoodsAndServicesTax `json:"goods_and_services_tax,omitempty"`
	AdvancedTradeOnlyVolume *float64             `json:"advanced_trade_only_volume,omitempty"` // Advanced Trade volume (non-inclusive of Pro) across assets, denoted in USD.
	AdvancedTradeOnlyFees   *float64             `json:"advanced_trade_only_fees,omitempty"`   // Advanced Trade fees (non-inclusive of Pro) across assets, denoted in USD.
	CoinbaseProVolume       *float64             `json:"coinbase_pro_volume,omitempty"`        // Coinbase Pro volume across assets, denoted in USD.
	CoinbaseProFees         *float64             `json:"coinbase_pro_fees,omitempty"`          // Coinbase Pro fees across assets, denoted in USD.
	TotalBalance            *string              `json:"total_balance,omitempty,omitempty"`    // Total balance across assets and products, which is comprised of the sum of spot, intx, and fcm, and denoted in USD.
}

//	Get a summary of transactions with fee tiers, total volume, and fees.
//
// https://docs.cdp.coinbase.com/coinbase-app/trade/reference/retailbrokerageapi_gettransactionsummary
func (s *FeesService) GetTransactionsSummary(
	ctx context.Context,
	options *GetTransactionsSummaryOptions,
) (*GetTransactionsSummaryResponse, error) {

	u := s.client.baseURL + "/api/v3/brokerage/transaction_summary"

	var summary GetTransactionsSummaryResponse

	err := s.client.get(ctx, u, options, &summary)
	if err != nil {
		err = fmt.Errorf("failed to fetch get transactions summary: %w", err)
	}

	return &summary, err
}
