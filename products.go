// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"time"
)

type ContractExpiryType string

const (
	ContractExpiryTypeUnknown   ContractExpiryType = "UNKNOWN_CONTRACT_EXPIRY_TYPE"
	ContractExpiryTypeExpiring  ContractExpiryType = "EXPIRING"
	ContractExpiryTypePerpetual ContractExpiryType = "PERPETUAL"
)

type ProductsService service

type BidAsk struct {
	Price *string `json:"price"`
	Size  *string `json:"size"`
}

type PriceBook struct {
	ProductID string     `json:"product_id"`
	Bids      []BidAsk   `json:"bids"`
	Asks      []BidAsk   `json:"asks"`
	Time      *time.Time `json:"time"`
}

type ProductType string

const (
	ProductTypeSpot ProductType = "SPOT"
	Future          ProductType = "FUTURE"
)

type RiskManagedBy string

const (
	RiskManagedByUnknown RiskManagedBy = "UNKNOWN_RISK_MANAGEMENT_TYPE"
	RiskManagedByFCM     RiskManagedBy = "MANAGED_BY_FCM"
	RiskManagedByVenue   RiskManagedBy = "MANAGED_BY_VENUE"
)

type FCMTradingSessionDetails struct {
	IsSessionOpen *bool      `json:"is_session_open"`
	OpenTime      *time.Time `json:"open_time"`
	CloseTime     *time.Time `json:"close_time"`
}

type FutureProductDetails struct {
	Venue                  *string             `json:"venue"`
	ContractCode           *string             `json:"contract_code"`
	ContractExpiry         *time.Time          `json:"contract_expiry"`
	ContractSize           *string             `json:"contract_size"`
	ContractRootUnit       *string             `json:"contract_root_unit"`
	GroupDescription       *string             `json:"group_description"` // Descriptive name for the product series, eg "Nano Bitcoin Futures".
	ContractExpiryTimezone *string             `json:"contract_expiry_timezone"`
	GroupShortDescription  *string             `json:"group_short_description"` // Short version of the group_description, eg "Nano BTC".
	RiskManagedBy          *RiskManagedBy      `json:"risk_managed_by"`
	ContractExpiryType     *ContractExpiryType `json:"contract_expiry_type"`
	PerpetualDetails       *struct {
		OpenInterest *string    `json:"open_interest"`
		FundingRate  *string    `json:"funding_rate"`
		FundingTime  *time.Time `json:"funding_time"`
	} `json:"perpetual_details"`
	ContractDisplayName *string `json:"contract_display_name"`
}

type Product struct {
	ID                            string                    `json:"product_id"`                   // The trading pair.
	Price                         string                    `json:"price"`                        // The current price for the product, in quote currency.
	PricePercentageChange24Hours  string                    `json:"price_percentage_change_24h"`  // The amount the price of the product has changed, in percent, in the last 24 hours.
	Volume24Hours                 string                    `json:"volume_24h"`                   // The trading volume for the product in the last 24 hours.
	VolumePercentageChange24Hours string                    `json:"volume_percentage_change_24h"` // The percentage amount the volume of the product has changed in the last 24 hours.
	BaseIncrement                 string                    `json:"base_increment"`               // Minimum amount base value can be increased or decreased at once.
	QuoteIncrement                string                    `json:"quote_increment"`              // Minimum amount quote value can be increased or decreased at once.
	QuoteMinimumSize              string                    `json:"quote_min_size"`               // Minimum size that can be represented of quote currency.
	QuoteMaximumSize              string                    `json:"quote_max_size"`               // Maximum size that can be represented of quote currency.
	BaseMinimimSize               string                    `json:"base_min_size"`                // Minimum size that can be represented of base currency.
	BaseMaximumSize               string                    `json:"base_max_size"`                // Maximum size that can be represented of base currency.
	BaseName                      string                    `json:"base_name"`                    // Name of the base currency.
	QuoteName                     string                    `json:"quote_name"`                   // Name of the quote currency.
	Watched                       bool                      `json:"watched"`                      // Whether or not the product is on the user's watchlist.
	IsDisabled                    bool                      `json:"is_disabled"`                  // Whether or not the product is disabled for trading.
	New                           bool                      `json:"new"`                          // Whether or not the product is 'new'.
	Status                        string                    `json:"status"`                       // Status of the product.
	CancelOnly                    bool                      `json:"cancel_only"`                  // Whether or not orders of the product can only be cancelled, not placed or edited.
	LimitOnly                     bool                      `json:"limit_only"`                   // Whether or not orders of the product can only be limit orders, not market orders.
	PostOnly                      bool                      `json:"post_only"`                    // Whether or not orders of the product can only be posted, not cancelled.
	TradingDisabled               bool                      `json:"trading_disabled"`             // Whether or not the product is disabled for trading for all market participants.
	AuctionMode                   bool                      `json:"auction_mode"`                 // Whether or not the product is in auction mode.
	Type                          *ProductType              `json:"product_type"`                 // Type of product.
	QuoteCurrencyID               *string                   `json:"quote_currency_id"`            // Symbol of the quote currency.
	BaseCurrencyID                *string                   `json:"base_currency_id"`             // Symbol of the base currency.
	FCMTradingSessionDetails      *FCMTradingSessionDetails `json:"fcm_trading_session_details"`
	MidMarketPrice                *string                   `json:"mid_market_price"`     // The current midpoint of the bid-ask spread, in quote currency.
	Alias                         *string                   `json:"alias"`                // Product id for the corresponding unified book.
	AliasTo                       []string                  `json:"alias_to"`             // Product ids that this product serves as an alias for.
	BaseDisplaySymbol             string                    `json:"base_display_symbol"`  // Symbol of the base display currency.
	QuoteDisplaySymbol            string                    `json:"quote_display_symbol"` // Symbol of the quote display currency.
	ViewOnly                      *bool                     `json:"view_only"`            // Whether or not the product is in view only mode.
	PriceIncrement                *string                   `json:"price_increment"`      // Minimum amount price can be increased or decreased at once.
	FutureProductDetails          *FutureProductDetails     `json:"future_product_details"`
}
