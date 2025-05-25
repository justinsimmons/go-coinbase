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
	ID                *uuid.UUID       `json:"uuid,omitempty"`                // Unique identifier for account.
	Name              *string          `json:"name,omitempty"`                // Name for the account.
	Currency          *string          `json:"currency,omitempty"`            // Currency symbol for the account.
	AvailableBalance  AvailableBalance `json:"available_balance"`             // Available balance of account.
	Default           *bool            `json:"default,omitempty"`             // Whether or not this account is the user's primary account.
	Active            *bool            `json:"active,omitempty"`              // Whether or not this account is active and okay to use.
	CreatedAt         *time.Time       `json:"created_at,omitempty"`          // Time at which this account was created.
	UpdatedAt         *time.Time       `json:"updated_at,omitempty"`          // Time at which this account was updated.
	DeletedAt         *time.Time       `json:"deleted_at,omitempty"`          // Time at which this account was deleted.
	Type              *AccountType     `json:"type,omitempty"`                // Type of account.
	Ready             *bool            `json:"ready,omitempty"`               // Whether or not this account is ready to trade.
	Hold              Hold             `json:"hold"`                          // Amount that is being held for pending transfers against the available balance.
	RetailPortfolioID *string          `json:"retail_portfolio_id,omitempty"` // The ID of the portfolio this account is associated with.
	Platform          *AccountPlatform `json:"platform,omitempty"`            // Platform indicates if the account is for spot (CONSUMER), US Derivatives (CFM_CONSUMER), or International Exchange (INTX).
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
	Cursor *string `json:"cursor,omitempty"`
	// Number of accounts returned
	Size *int32 `json:"size,omitempty"`
}

// TradeIncentiveMetadata contains promo or incentive metadata for a convert
// request.
type TradeIncentiveMetadata struct {
	// UserIncentiveID is the ID of the user incentive
	UserIncentiveID string `json:"user_incentive_id,omitempty"`
	// CodeVal is the promo code for waiving fees.
	CodeVal string `json:"code_val,omitempty"`
}

// Status of a trade.
//
//go:generate enumer -type=AccountPlatform -transform=snake-upper -json -txt
type TradeStatus byte

const (
	TradeStatusUnspecified TradeStatus = iota
	TradeStatusCreated
	TradeStatusStarted
	TradeStatusCompleted
	TradeStatusCanceled
)

// Amount represents a monetary value in a specific currency.
type Amount struct {
	Value    *string `json:"value,omitempty"`    // The amount as a string (to preserve precision).
	Currency *string `json:"currency,omitempty"` // The ISO currency code (e.g. USD, BTC).
}

// Link.
type Link struct {
	Text *string `json:"text"`
	URL  *string `json:"url"`
}

// Disclosure contains additional information or fine print about a fee.
type Disclosure struct {
	Title       *string `json:"title,omitempty"`       // Title of the disclosure.
	Description *string `json:"description,omitempty"` // Description or explanation.
	Link        *Link   `json:"link,omitempty"`        // Optional hyperlink with more information.
}

// Source of a waived free.
//
//go:generate enumer -type=WaivedFeeSource -transform=snake-upper -json -txt
type WaivedFeeSource byte

const (
	WaivedFeeSourceUnspecified WaivedFeeSource = iota
	WaivedFeeSourceTradeIncentive
	WaivedFeeSourceCoinbaseOne
	WaivedFeeSourceFeeReductionBadge
	WaivedFeeSourceTradePromotion
)

// WaivedDetail describes the waived portion of a fee, if applicable.
// type WaivedDetail struct {
// 	Amount *Amount          `json:"amount"` // The waived amount.
// 	Source *WaivedFeeSource `json:"source"` // Source of the waived fee.
// }
//
// // Fee represents a single fee item applied to the trade.
// type Fee struct {
// 	Title        *string       `json:"title,omitempty"`          // Short name for the fee.
// 	Description  *string       `json:"description,omitempty"`    // Detailed description of the fee.
// 	Amount       *Amount       `json:"amount,omitempty"`         // The fee amount.
// 	Label        *string       `json:"label,omitempty"`          // A short label for the fee.
// 	Disclosure   *Disclosure   `json:"disclosure,omitempty"`     // Optional disclosure information
// 	WaivedDetail *WaivedDetail `json:"waived_details,omitempty"` // Optional info on any waived fee amount.
// }
//
// // Source represents the origin of funds or the payment method used for the trade.
// type Source struct {
// 	Type              string             `json:"type"`                         // Source type (e.g. COINBASE_ACCOUNT, BLOCKCHAIN_ADDRESS)
// 	BlockchainAddress *BlockchainAddress `json:"blockchain_address,omitempty"` // Blockchain address source (if used)
// 	CoinbaseAccount   *CoinbaseAccount   `json:"coinbase_account,omitempty"`   // Coinbase account source (if used)
// 	// Additional sources like fedwire, swift, paypal can be added here as needed
// }
//
// type ConvertTrade struct {
// 	ID                string      `json:"id,omitempty"`                // The unique ID for the trade.
// 	Status            TradeStatus `json:"status"`                      // The status of the trade (e.g. TRADE_STATUS_CREATED).
// 	UserEnteredAmount Amount      `json:"user_entered_amount"`         // The amount originally entered by the user.
// 	Amount            Amount      `json:"amount"`                      // The computed amount to be converted.
// 	Subtotal          Amount      `json:"subtotal"`                    // The subtotal amount before fees.
// 	Total             Amount      `json:"total"`                       // The total amount including fees.
// 	Fees              []Fee       `json:"fees"`                        // List of individual fees applied to the trade.
// 	TotalFee          Fee         `json:"total_fee"`                   // The total of all fees applied to the trade.
// 	Source            *Source     `json:"source,omitempty"`            // The source of funds or payment (e.g. account, wallet).
// 	// Target Target
// 	Network           string      `json:"network,omitempty"`           // The payment network used (if applicable).
// 	PaymentMethodID   string      `json:"payment_method_id,omitempty"` // ID of the payment method used (if applicable).
// }

// Product type.
//
//go:generate enumer -type=ProductType -transform=snake-upper -trimprefix=ProductType -json -text
type ProductType byte

const (
	ProductTypeUnknownProductType ProductType = iota
	ProductTypeSpot
	ProductTypeFuture
)

// Contract expiry type.
//
//go:generate enumer -type=ContractExpiryType -transform=snake-upper -trimprefix=ContractExpiryType -json -text
type ContractExpiryType byte

const (
	ContractExpiryTypeUnknonContractExpiryType ContractExpiryType = iota
	ContractExpiryTypeExpiring
	ContractExpiryTypePerpetual
)

// Venue for a product..
//
//go:generate enumer -type=ProductVenue -transform=snake-upper -trimprefix=ProductVenue -json -text
type ProductVenue byte

const (
	ProductVenueUnknownProductVenue ProductVenue = iota
	ProductVenueCbe
	ProductVenueFcm
	ProductVenueIntx
)

// Description of maker and taker rates across all applicable fee tiers.
type FeeTier struct {
	PricingTier  *string `json:"pricing_tier,omitempty"`   // Pricing tier for user, determined by notional (USD) volume.
	USDFrom      *string `json:"usd_from,omitempty"`       // Lower bound (inclusive) of pricing tier in notional volume.
	USDTo        *string `json:"usd_to,omitempty"`         // Upper bound (exclusive) of pricing tier in notional volume.
	TakerFeeRate *string `json:"taker_fee_rate,omitempty"` // Taker fee rate, applied if the order takes liquidity.
	MakerFeeRate *string `json:"maker_fee_rate,omitempty"` // Maker fee rate, applied if the order creates liquidity.
	AOPFrom      *string `json:"aop_from,omitempty"`       // Lower bound (inclusive) of pricing tier in usd of total assets on platform.
	AOPTO        *string `json:"aop_to,omitempty"`         // Upper bound (exclusive) of pricing tier in usd of total assets on platform.
	PerpsVolFrom *string `json:"perps_vol_from,omitempty"` // Lower bound (inclusive) of pricing tier in notional volume for perpetual contracts.
	PerpsVolTo   *string `json:"perps_vol_to,omitempty"`   // Upper bound (exclusive) of pricing tier in notional volume for perpetual contracts.
}

// Tyype of goods and services tax.
type GoodsAndServicesType byte

const (
	GoodsAndServicesTypeInclusive GoodsAndServicesType = iota
	GoodsAndServicesTypeExclusive
)

// Goods and services tax information.
type GoodsAndServicesTax struct {
	Rate *string               `json:"rate,omitempty"`
	Type *GoodsAndServicesType `json:"type,omitempty"`
}

// Margin rate.
type MarginRate struct {
	Value *string `json:"value,omitempty"` // String representation allows for unlimited precision.
}
