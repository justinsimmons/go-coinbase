// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"time"

	"github.com/google/uuid"
)

// Type of coinbase account.
//
//go:generate enumer -type=AccountType -transform=snake-upper -json -txt
type AccountType byte

const (
	AccountTypeUnspecified AccountType = iota
	AccountTypeCrypto
	AccountTypeFiat
	AccountTypeVault
	AccountTypePerpFutures
)

// Coinbase platform the account is on.
//
//go:generate enumer -type=AccountPlatform -transform=snake-upper -json -txt
type AccountPlatform byte

const (
	PlatformUnspecified AccountPlatform = iota // Unspecified or unknown.
	PlatformConsumer                           // SPOT consumer.
	PlatformCfmConsumer                        // US Derivatives.
	PlatformIntx                               // International Exchange.
)

// Available balance belonging to coinbase account.
type AvailableBalance struct {
	Value    string `json:"value"`    // Amount of currency that this object represents.
	Currency string `json:"currency"` // Denomination of the currency.
}

// Amount that is being held for pending transfers against the available
// balance.
type Hold struct {
	Value    string `json:"value"`    // Amount of currency that this object represents.
	Currency string `json:"currency"` // Denomination of the currency.
}

// Coinbase account data.
type Account struct {
	ID                *uuid.UUID       `json:"uuid"`                // Unique identifier for account.
	Name              *string          `json:"name"`                // Name for the account.
	Currency          *string          `json:"currency"`            // Currency symbol for the account.
	AvailableBalance  AvailableBalance `json:"available_balance"`   // Available balance of account.
	Default           *bool            `json:"default"`             // Whether or not this account is the user's primary account.
	Active            *bool            `json:"active"`              // Whether or not this account is active and okay to use.
	CreatedAt         *time.Time       `json:"created_at"`          // Time at which this account was created.
	UpdatedAt         *time.Time       `json:"updated_at"`          // Time at which this account was updated.
	DeletedAt         *time.Time       `json:"deleted_at"`          // Time at which this account was deleted.
	Type              *AccountType     `json:"type"`                // Type of account.
	Ready             *bool            `json:"ready"`               // Whether or not this account is ready to trade.
	Hold              Hold             `json:"hold"`                // Amount that is being held for pending transfers against the available balance.
	RetailPortfolioID *string          `json:"retail_portfolio_id"` // The ID of the portfolio this account is associated with.
	Platform          *AccountPlatform `json:"platform"`            // Platform indicates if the account is for spot (CONSUMER), US Derivatives (CFM_CONSUMER), or International Exchange (INTX).
}

type PaginationOptions struct {
	// The number of accounts to display per page. By default, displays 49
	// (max 250). If has_next is true, additional pages of accounts are
	// available to be fetched. Use the cursor parameter to start on a
	// specified page.
	Limit int32 `url:"limit,omitempty"`
	// For paginated responses, returns all responses that come after this
	// value.
	Cursor *string `url:"cursor,omitempty"`
}

// Pagination response from the Coinbase API.
type PaginatedResponse struct {
	// Whether there are additional pages for this query.
	HasNext bool `json:"has_next"`
	// For paginated responses, returns all responses that come after this
	// value.
	Cursor *string `json:"cursor"`
	// Number of accounts returned
	Size *int32 `json:"size"`
}
