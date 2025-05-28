// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

// AccountsService handles communication with the accounts related methods of
// the Coinbase Advanced Trade API.
type AccountsService service

// ConvertsService handles communication with the convert trades related
// methods of the Coinbase Advanced Trade API.
type ConvertsService service

// DataApiService handles communication with the data API related
// methods of the Coinbase Advanced Trade API.
type DataApiService service

// FeesService handles communication with the fees related methods of the
// Coinbase Advanced Trade API.
type FeesService service

// FuturesService handles communication with the futures related methods of the
// Coinbase Advanced Trade API.
// Futures vs Spot Accounts
//
//	Futures and spot balances are held in different accounts. Cash is always
//
// deposited into your Coinbase Inc. (CBI) spot account. You can only acquire
// spot assets with funds in your spot account.
//
// Treatment of Cash
//
//	Cash is automatically transferred to your Coinbase Financial Markets (CFM)
//
// futures account to satisfy margin requirements. Automatic transfers are
// only from CBI spot accounts to CFM futures accounts. You can transfer cash
// that isn't being used to margin or maintain futures positions into your
// CBI spot account (to trade spot assets or to withdraw) with Schedule
// Futures Sweep. Funds held in a CBI spot account do not receive the
// preferential treatment given to funds held in a regulated futures account,
// pursuant to CFTC's regulations and the U.S. Bankruptcy Code.
type FuturesService service

// OrdersService handles communication with the orders related methods
// of the Coinbase Advanced Trade API.
type OrdersService service

// PerpetualsService handles communication with the perpetuals related methods
// of the Coinbase Advanced Trade API.
type PerpetualsService service

// PaymentMethodsService handles communication with the payment method related
// methods of the Coinbase Advanced Trade API.
type PaymentMethodsService service
