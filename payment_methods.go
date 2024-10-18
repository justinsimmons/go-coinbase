// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import "time"

type PaymentMethodsService service

type PaymentMethod struct {
	ID            string    `json:"id"`             // Unique identifier for the payment method.
	Type          string    `json:"type"`           // The payment method type.
	Name          string    `json:"name"`           // Name for the payment method.
	Currency      string    `json:"currency"`       // Currency symbol for the payment method.
	Verified      bool      `json:"verified"`       // The verified status of the payment method.
	AllowBuy      bool      `json:"allow_buy"`      // Whether or not this payment method can perform buys.
	AllowSell     bool      `json:"allow_sell"`     // Whether or not this payment method can perform sells.
	AllowDeposit  bool      `json:"allow_deposit"`  // Whether or not this payment method can perform deposits.
	AllowWithdraw bool      `json:"allow_withdraw"` // Whether or not this payment method can perform withdrawals.
	CreatedAt     time.Time `json:"created_at"`     // Time at which this payment method was created.
	UpdatedAt     time.Time `json:"updated_at"`     // Time at which this payment method was updated.
}
