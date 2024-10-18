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

type PortfolioBalances struct {
	TotalBalance               Funds `json:"total_balance"`                 // Represents a monetary amount.
	TotalFuturesBalance        Funds `json:"total_futures_balance"`         // Represents a monetary amount.
	TotalCashEquivalentBalance Funds `json:"total_cash_equivalent_balance"` // Represents a monetary amount.
	TotalCryptoBalance         Funds `json:"total_crypto_balance"`          // Represents a monetary amount.
	FuturesUnrealizedPNL       Funds `json:"futures_unrealized_pnl"`        // Represents a monetary amount.
	PerpUnrealizedPNL          Funds `json:"perp_unrealized_pnl"`           // Represents a monetary amount.

}

type SpotPosition struct {
	Asset                *string    `json:"asset"`
	AccountUUID          *uuid.UUID `json:"account_uuid"`
	TotalBalanceFiat     *float32   `json:"total_balance_fiat"`
	TotalBalanceCrypto   *float32   `json:"total_balance_crypto"`
	AvailableToTradeFiat *float32   `json:"available_to_trade_fiat"`
	Allocation           *float32   `json:"allocation"`
	OneDayChange         *float32   `json:"one_day_change"`
	CostBasis            Funds      `json:"cost_basis"` // Represents a monetary amount.
	AssetImageURL        *string    `json:"asset_img_url"`
	IsCash               *bool      `json:"is_cash"`
}

type Currency struct {
	UserNativeCurrency Funds `json:"userNativeCurrency"` // Represents a monetary amount.
	RawCurrency        Funds `json:"rawCurrency"`        // Represents a monetary amount.
}

type PerpPosition struct {
	ProductID        *string              `json:""`
	ProductUUID      *uuid.UUID           `json:""`
	Symbol           *string              `json:""`
	AssetImageURL    *string              `json:"asset_img_url"`
	VWAP             *Currency            `json:""`
	PositionSide     *FuturesPositionSide `json:""`
	NetSize          *string              `json:""`
	BuyOrderSize     *string              `json:""`
	SellOrderSize    *string              `json:""`
	IMContribution   *string              `json:""`
	UnrealizedPNL    *Currency
	MarkPrice        *Currency
	LiquidationPrice *Currency
	Leverage         *string
	IMNotional       *Currency
	MMNotional       *Currency
	PositionNotional *Currency
}

type PortfolioBreakdown struct {
	Portfolio     *Portfolio         `json:""`
	Balances      *PortfolioBalances `json:""`
	SpotPositions []SpotPosition     `json:"spot_positions"`
	PerpPositions []PerpPosition     `json:""`
}

type portfolioBreakdownResponse struct {
	Breakdown PortfolioBreakdown `json:"breakdown"`
}

// GetPortfolioBreakdown gets the breakdown of a portfolio by portfolio ID.
func (s *PortfoliosService) GetPortfolioBreakdown(ctx context.Context, id string) (*PortfolioBreakdown, error) {
	var breakdownResp portfolioBreakdownResponse

	err := s.client.get(ctx, fmt.Sprintf("%s/api/v3/brokerage/portfolios/%s", s.client.baseURL, id), nil, &breakdownResp)
	if err != nil {
		return nil, fmt.Errorf("failed to get portfolio breakdown for '%s': %w", id, err)
	}

	return &breakdownResp.Breakdown, err
}
