// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import "time"

type FuturesSide string

const (
	FuturesSideUnknown FuturesSide = "UNKNOWN"
	FuturesSideLong    FuturesSide = "LONG"
	FuturesSideShort   FuturesSide = "SHORT"
)

// Interface for interacting with Fututres APIs.
//
// Futures vs Spot Accounts
//
//	Futures and spot balances are held in different accounts. Cash is always deposited into your Coinbase Inc. (CBI) spot account. You can only acquire spot assets with funds in your spot account.
//
// Treatment of Cash
//
//	Cash is automatically transferred to your Coinbase Financial Markets (CFM) futures account to satisfy margin requirements. Automatic transfers are only from CBI spot accounts to CFM futures accounts.
//	You can transfer cash that isn't being used to margin or maintain futures positions into your CBI spot account (to trade spot assets or to withdraw) with Schedule Futures Sweep.
//	Funds held in a CBI spot account do not receive the preferential treatment given to funds held in a regulated futures account, pursuant to CFTC's regulations and the U.S. Bankruptcy Code.
type FuturesService service

type FuturesPosition struct {
	ProductID         *string      `json:"product_id"`          // The ID of the CFM futures product.
	ExpirationTime    *time.Time   `json:"expiration_time"`     // The expiry of your position.
	Side              *FuturesSide `json:"side"`                // The side of your position.
	NumberOfContracts *string      `json:"number_of_contracts"` // The size of your position in contracts.
	CurrentPrice      *string      `json:"current_price"`       // The current price of the product.
	AverageEntryPrice *string      `json:"avg_entry_price"`     // The average entry price at which you entered your current position.
	UnrealizedPNL     *string      `json:"unrealized_pnl"`      // Your current unrealized PnL for your position.
	DailyRealizedPNL  *string      `json:"daily_realized_pnl"`  // Your realized PnL from your trades in this product on current trade date.
}
