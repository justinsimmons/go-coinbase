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

type BalanceSummary struct {
	FuturesBuyingPower          Funds `json:"futures_buying_power"`          // The amount of your cash balance that is available to trade CFM futures.
	TotalBalance                Funds `json:"total_usd_balance"`             // Aggregate USD maintained across your CFTC-regulated futures account and your Coinbase Inc. spot account
	CBIBalance                  Funds `json:"cbi_usd_balance"`               // USD maintained in your Coinbase Inc. spot account.
	CFMBalance                  Funds `json:"cfm_usd_balance"`               // USD maintained in your CFTC-regulated futures account. Funds held in your futures account are not available to trade spot.
	TotalOpenOrdersHoldAmmount  Funds `json:"total_open_orders_hold_amount"` // Your total balance on hold for spot and futures open orders.
	UnrealizedPNL               Funds `json:"unrealized_pnl"`                // Your current unrealized PnL across all open positions.
	DailyRealizedPNL            Funds `json:"daily_realized_pnl"`            // Your realized PnL from the current trade date. May include profit or loss from positions you’ve closed on the current trade date.
	InitialMargin               Funds `json:"initial_margin"`                // Margin required to initiate futures positions. Once futures orders are placed, these funds cannot be used to trade spot. The actual amount of funds necessary to support executed futures orders will be moved to your futures account.
	AvailableMargin             Funds `json:"available_margin"`              // Funds available to meet your anticipated margin requirement. This includes your CBI spot USD, CFM futures USD, and Futures PnL, less any holds for open spot or futures orders.
	LiquidationThreshold        Funds `json:"liquidation_threshold"`         // When your available funds for collateral drop to the liquidation threshold, some or all of your futures positions will be liquidated.
	LiquidationBufferAmount     Funds `json:"liquidation_buffer_amount"`     // Funds available in excess of the liquidation threshold, calculated as available margin minus liquidation threshold. If your liquidation buffer amount reaches 0, your futures positions and/or open orders will be liquidated as necessary.
	LiquidationBufferPercentage Funds `json:"liquidation_buffer_percentage"` // Funds available in excess of the liquidation threshold expressed as a percentage. If your liquidation buffer percentage reaches 0%, your futures positions and/or open orders will be liquidated as necessary.
}

// GetBalanceSummary gets information on your balances related to Coinbase Financial Markets (CFM) futures trading.
type getBalanceSummaryResponse struct {
	BalanceSummary *BalanceSummary `json:"balance_summary"`
}

func (s *FuturesService) GetBalanceSummary(ctx context.Context) (*BalanceSummary, error) {
	var resp getBalanceSummaryResponse

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/cfm/balance_summary", nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance summary: %w", err)
	}

	return resp.BalanceSummary, nil
}
