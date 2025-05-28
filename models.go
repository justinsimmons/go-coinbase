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
//go:generate enumer -type=AccountType -transform=snake-upper -json -text
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
//go:generate enumer -type=AccountPlatform -transform=snake-upper -json -text
type AccountPlatform byte

const (
	PlatformUnspecified AccountPlatform = iota // Unspecified or unknown.
	PlatformConsumer                           // SPOT consumer.
	PlatformCfmConsumer                        // US Derivatives.
	PlatformIntx                               // International Exchange.
)

// Available balance belonging to coinbase account.
// TODO:: Rename to something more generic.
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
//go:generate enumer -type=AccountPlatform -transform=snake-upper -json -text
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
//go:generate enumer -type=WaivedFeeSource -transform=snake-upper -json -text
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

// Margin profile type.
//
//go:generate enumer -type=AccountType -transform=snake-upper -json -text
type MarginProfileType byte

const (
	MarginProfileTypeUnspecified MarginProfileType = iota
	MarginProfileTypeRetailRegular
	MarginProfileTypeRetailIntradayMargin1 // TODO: Will this be Margin1 or Margin_1
)

// Margin window type.
//
//go:generate enumer -type=MarginWindowType -transform=snake-upper -json -text
type MarginWindowType byte

const (
	MarginWindowTypeUnspecified MarginWindowType = iota
	MarginWindowTypeOvernight
	MarginWindowTypeWeekend
	MarginWindowTypeIntraday
	MarginWindowTypeTransition
)

type MarginWindow struct {
	Type    *MarginWindowType `json:"margin_window_type,omitempty"` // Your margin window.
	EndTime *time.Time        `json:"end_time,omitempty"`           // The end time of the margin window.
}

// Margin level types.
//
//go:generate enumer -type=MarginLevelType -transform=snake-upper -json -text
type MarginLevelType byte

const (
	MarginLevelTypeUnspecified MarginLevelType = iota
	MarginLevelTypeBase
	MarginLevelTypeWarning
	MarginLevelTypeDanger
	MarginLevelTypeLiquidation
)

// Margin window measure.
type MarginWindowMeasure struct {
	MarginWindowType   *MarginWindowType `json:"margin_window_type,omitempty"` // Your margin window.
	MarginLevel        *MarginLevelType  `json:"margin_level,omitempty"`       // Margin level for liquidation purposes.
	InitialMargin      *string           `json:"initial_margin,omitempty"`
	MaintenanceMargin  *string           `json:"maintenance_margin,omitempty"`
	LiquidationBuffer  *string           `json:"liquidation_buffer,omitempty"`
	TotalHold          *string           `json:"total_hold,omitempty"`
	FuturesBuyingPower *string           `json:"futures_buying_power,omitempty"` // The amount of your cash balance that is available to trade CFM futures.
}

// Balance summary.
type BalanceSummary struct {
	FuturesBuyingPower           *AvailableBalance    `json:"futures_buying_power,omitempty"`          // The amount of your cash balance that is available to trade CFM futures.
	TotalUsdBalance              *AvailableBalance    `json:"total_usd_balance,omitempty"`             // Aggregate USD maintained across your CFTC-regulated futures account and your Coinbase Inc. spot account.
	CbiUsdBalance                *AvailableBalance    `json:"cbi_usd_balance,omitempty"`               // USD maintained in your Coinbase Inc. spot account.
	CfmUsdBalance                *AvailableBalance    `json:"cfm_usd_balance,omitempty"`               // USD maintained in your CFTC-regulated futures account. Funds held in your futures account are not available to trade spot.
	TotalOpenOrdersHoldAmmount   *AvailableBalance    `json:"total_open_orders_hold_amount,omitempty"` // Your total balance on hold for spot and futures open orders.
	UnrealizedPNL                *AvailableBalance    `json:"unrealized_pnl,omitempty"`                // Your current unrealized PnL across all open positions.
	DailyRealizedPNL             *AvailableBalance    `json:"daily_realized_pnl,omitempty"`            // Your realized PnL from the current trade date. May include profit or loss from positions you’ve closed on the current trade date.
	InitialMargin                *AvailableBalance    `json:"initial_margin,omitempty"`                // Margin required to initiate futures positions. Once futures orders are placed, these funds cannot be used to trade spot. The actual amount of funds necessary to support executed futures orders will be moved to your futures account.
	AvailableMargin              *AvailableBalance    `json:"available_margin,omitempty"`              // Funds available to meet your anticipated margin requirement. This includes your CBI spot USD, CFM futures USD, and Futures PnL, less any holds for open spot or futures orders.
	LiquidationThreshold         *AvailableBalance    `json:"liquidation_threshold,omitempty"`         // When your available funds for collateral drop to the liquidation threshold, some or all of your futures positions will be liquidated.
	LiquidationBufferAmount      *AvailableBalance    `json:"liquidation_buffer_amount,omitempty"`     // Funds available in excess of the liquidation threshold, calculated as available margin minus liquidation threshold. If your liquidation buffer amount reaches 0, your futures positions and/or open orders will be liquidated as necessary.
	LiquidationBufferPercentage  *string              `json:"liquidation_buffer_percentage,omitempty"` // Funds available in excess of the liquidation threshold expressed as a percentage. If your liquidation buffer percentage reaches 0%, your futures positions and/or open orders will be liquidated as necessary.
	IntradayMarginWindowMeasure  *MarginWindowMeasure `json:"intraday_margin_window_measure,omitempty"`
	OvernightMarginWindowMeasure *MarginWindowMeasure `json:"overnight_margin_window_measure,omitempty"`
}

// The side of a futures trading position.
//
//go:generate enumer -type=FuturesSide -transform=snake-upper -trimprefix=FuturesSide -json -text
type FuturesSide byte

const (
	FuturesSideUnknown FuturesSide = iota
	FuturesSideLong
	FuturesSideShort
)

// Futures trading position.
type FuturesPosition struct {
	ProductID         *string      `json:"product_id,omitempty"`          // The ticker symbol (e.g. 'BIT-28JUL23-CDE').
	ExpirationTime    *time.Time   `json:"expiration_time,omitempty"`     // The expiry of your position.
	Side              *FuturesSide `json:"side,omitempty"`                // The side of your position.
	NumberOfContracts *string      `json:"number_of_contracts,omitempty"` // The size of your position in contracts.
	CurrentPrice      *string      `json:"current_price,omitempty"`       // The current price of the product.
	AverageEntryPrice *string      `json:"avg_entry_price,omitempty"`     // The average entry price at which you entered your current position.
	UnrealizedPNL     *string      `json:"unrealized_pnl,omitempty"`      // Your current unrealized PnL for your position.
	DailyRealizedPNL  *string      `json:"daily_realized_pnl,omitempty"`  // Your realized PnL from your trades in this product on current trade date.
}

// Intraday margin setting.
//
//go:generate enumer -type=IntradayMarginSetting -transform=snake-upper -json -text
type IntradayMarginSetting byte

const (
	IntradayMarginSettingUnspecified IntradayMarginSetting = iota
	IntradayMarginSettingStandard
	IntradayMarginSettingIntraday
)

// Futures sweep status.
//
//go:generate enumer -type=SweepStatus -transform=snake-upper -trimprefix=SweepStatus -json -text
type SweepStatus byte

const (
	SweepStatusUnknownFcmSweepStatus SweepStatus = iota
	SweepStatusPending
	SweepStatusProcessing
)

// Futures sweep.
type FuturesSweep struct {
	ID              *string           `json:"id,omitempty"`               // The ID of the sweep request scheduled.
	RequestedAmount *AvailableBalance `json:"requested_amount,omitempty"` //The requested sweep amount.
	ShouldSweepAll  bool              `json:"should_sweep_all,omitempty"` // True if the request was to sweep all available funds from your CFM futures account.
	Status          *SweepStatus      `json:"status,omitempty"`           // A pending sweep is a sweep that has not started processing and can be cancelled. A processing sweep is a sweep that is currently being processed and cannot be cancelled.
	ScheduledTime   *time.Time        `json:"scheduled_time,omitempty"`   // The timestamp at which the sweep request was submitted.
}

// Reason for failure to cancel order.
//
//go:generate enumer -type=CancelOrderFailureReason -transform=snake-upper -trimprefix=CancelOrderFailureReason -json -text
type CancelOrderFailureReason byte

const (
	CancelOrderFailureReasonUnknownCancelFailureReason CancelOrderFailureReason = iota
	CancelOrderFailureReasonInvalidCancelRequest
	CancelOrderFailureReasonUnknownCancelOrder
	CancelOrderFailureReasonCommanderRejectedCancelOrder
	CancelOrderFailureReasonDuplicateCancelRequest
	CancelOrderFailureReasonInvalidCancelProductId
	CancelOrderFailureReasonInvalidCancelFcmTradingSession
	CancelOrderFailureReasonNotAllowedToCancel
	CancelOrderFailureReasonOrderIsFullyFilled
	CancelOrderFailureReasonOrderIsBeingReplaced
)

// Result when order is attempted to be canceled.
type CancelOrderResult struct {
	Success       bool                     `json:"success"`        // Whether the cancel request was submitted successfully.
	FailureReason CancelOrderFailureReason `json:"failure_reason"` // The reason the cancel request did not get submitted.
	OrderID       string                   `json:"order_id"`       // The IDs of order cancel request was initiated for.
}

// Side of a trade.
//
//go:generate enumer -type=Side -transform=snake-upper -trimprefix=Side -json -text
type Side byte

const (
	SideBuy Side = iota
	SideSell
)

type OrderSuccessMetadata struct {
	OrderID       string  `json:"order_id"`                  // The ID of the order.
	ProductID     *string `json:"product_id,omitempty"`      // The trading pair (e.g. 'BTC-USD').
	Side          *Side   `json:"side,omitempty"`            // The side of the market that the order is on (e.g. 'BUY', 'SELL').
	ClientOrderID *string `json:"client_order_id,omitempty"` // The unique ID provided for the order (used for identification purposes).
}

// Order failure reason.
//
//go:generate enumer -type=OrderFailureReason -transform=snake-upper -trimprefix=OrderFailureReason -json -text
type OrderFailureReason byte

const (
	OrderFailureReasonUnknownFailureReason OrderFailureReason = iota
	OrderFailureReasonUnsupportedOrderConfiguration
	OrderFailureReasonInvalidSide
	OrderFailureReasonInvalidProductID
	OrderFailureReasonInvalidSizePrecision
	OrderFailureReasonInvalidPricePrecision
	OrderFailureReasonInsufficientFund
	OrderFailureReasonInvalidLedgerBalance
	OrderFailureReasonOrderEntryDisabled
	OrderFailureReasonIneligiblePair
	OrderFailureReasonInvalidLimitPricePostOnly
	OrderFailureReasonInvalidLimitPrice
	OrderFailureReasonInvalidNoLiquidity
	OrderFailureReasonInvalidRequest
	OrderFailureReasonCommanderRejectedNewOrder
	OrderFailureReasonInsufficientFunds
	OrderFailureReasonInLiquidation
	OrderFailureReasonInvalidMarginType
	OrderFailureReasonInvalidLeverage
	OrderFailureReasonUntradableProduct
	OrderFailureReasonInvalidFcmTradingSession
	OrderFailureReasonGeofencingRestriction
	OrderFailureReasonQuoteSize
	OrderFailureReasonQuoteSizeNotAllowedForBracket
	OrderFailureReasonInvalidBracketPrices
	OrderFailureReasonMissingMarketTradeData
	OrderFailureReasonInvalidBracketLimitPrice
	OrderFailureReasonInvalidBracketStopTriggerPrice
	OrderFailureReasonBracketLimitPriceOutOfBounds
	OrderFailureReasonStopTriggerPriceOutOfBounds
	OrderFailureReasonBracketOrderNotSupported
	OrderFailureReasonFokDisabled
	OrderFailureReasonFokOnlyAllowedOnLimitOrders
	OrderFailureReasonPostOnlyNotAllowedWithFok
	OrderFailureReasonUboHighLeverageQuantityBreached
	OrderFailureReasonEndTimeTooFarInFuture
	OrderFailureReasonLimitPriceTooFarFromMarket
	OrderFailureReasonOpenBracketOrders
	OrderFailureReasonFuturesAfterHourInvalidOrderType
	OrderFailureReasonFuturesAfterHourInvalidTimeInForce
	OrderFailureReasonInvalidAttachedTakeProfitPrice
	OrderFailureReasonInvalidAttachedStopLossPrice
	OrderFailureReasonInvalidAttachedTakeProfitPricePrecision
	OrderFailureReasonInvalidAttachedStopLossPricePrecision
	OrderFailureReasonInvalidAttachedTakeProfitPriceOutOfBounds
	OrderFailureReasonInvalidAttachedStopLossPriceOutOfBounds
	OrderFailureReasonInvalidAttachedTakeProfitPriceExceedsMaxDistance
	OrderFailureReasonInvalidAttachedTakeProfitSizeBelowMin
	OrderFailureReasonAttachedOrderSizeMustBeNil
)

// PreviewFailureReason represents the failure reason for a preview.
//
//go:generate enumer -type=PreviewFailureReason -transform=snake-upper -trimprefix=PreviewFailureReason -json -text
type PreviewFailureReason byte

const (
	PreviewFailureReasonUnknown PreviewFailureReason = iota
	PreviewFailureReasonPreviewMissingCommissionRate
	PreviewFailureReasonPreviewInvalidSide
	PreviewFailureReasonPreviewInvalidOrderConfig
	PreviewFailureReasonPreviewInvalidProductID
	PreviewFailureReasonPreviewInvalidSizePrecision
	PreviewFailureReasonPreviewInvalidPricePrecision
	PreviewFailureReasonPreviewMissingProductPriceBook
	PreviewFailureReasonPreviewInvalidLedgerBalance
	PreviewFailureReasonPreviewInsufficientLedgerBalance
	PreviewFailureReasonPreviewInvalidLimitPricePostOnly
	PreviewFailureReasonPreviewInvalidLimitPrice
	PreviewFailureReasonPreviewInvalidNoLiquidity
	PreviewFailureReasonPreviewInsufficientFund
	PreviewFailureReasonPreviewInvalidCommissionConfiguration
	PreviewFailureReasonPreviewInvalidStopPrice
	PreviewFailureReasonPreviewInvalidBaseSizeTooLarge
	PreviewFailureReasonPreviewInvalidBaseSizeTooSmall
	PreviewFailureReasonPreviewInvalidQuoteSizePrecision
	PreviewFailureReasonPreviewInvalidQuoteSizeTooLarge
	PreviewFailureReasonPreviewInvalidPriceTooLarge
	PreviewFailureReasonPreviewInvalidQuoteSizeTooSmall
	PreviewFailureReasonPreviewInsufficientFundsForFutures
	PreviewFailureReasonPreviewBreachedPriceLimit
	PreviewFailureReasonPreviewBreachedAccountPositionLimit
	PreviewFailureReasonPreviewBreachedCompanyPositionLimit
	PreviewFailureReasonPreviewInvalidMarginHealth
	PreviewFailureReasonPreviewRiskProxyFailure
	PreviewFailureReasonPreviewUntradableFcmAccountStatus
	PreviewFailureReasonPreviewInLiquidation
	PreviewFailureReasonPreviewInvalidMarginType
	PreviewFailureReasonPreviewInvalidLeverage
	PreviewFailureReasonPreviewUntradableProduct
	PreviewFailureReasonPreviewInvalidFcmTradingSession
	PreviewFailureReasonPreviewNotAllowedByMarketState
	PreviewFailureReasonPreviewBreachedOpenInterestLimit
	PreviewFailureReasonPreviewGeofencingRestriction
	PreviewFailureReasonPreviewInvalidEndTime
	PreviewFailureReasonPreviewOppositeMarginTypeExists
	PreviewFailureReasonPreviewQuoteSizeNotAllowedForBracket
	PreviewFailureReasonPreviewInvalidBracketPrices
	PreviewFailureReasonPreviewMissingMarketTradeData
	PreviewFailureReasonPreviewInvalidBracketLimitPrice
	PreviewFailureReasonPreviewInvalidBracketStopTriggerPrice
	PreviewFailureReasonPreviewBracketLimitPriceOutOfBounds
	PreviewFailureReasonPreviewStopTriggerPriceOutOfBounds
	PreviewFailureReasonPreviewBracketOrderNotSupported
	PreviewFailureReasonPreviewInvalidStopPricePrecision
	PreviewFailureReasonPreviewStopPriceAboveLimitPrice
	PreviewFailureReasonPreviewStopPriceBelowLimitPrice
	PreviewFailureReasonPreviewStopPriceAboveLastTradePrice
	PreviewFailureReasonPreviewStopPriceBelowLastTradePrice
	PreviewFailureReasonPreviewFokDisabled
	PreviewFailureReasonPreviewFokOnlyAllowedOnLimitOrders
	PreviewFailureReasonPreviewPostOnlyNotAllowedWithFok
	PreviewFailureReasonPreviewUboHighLeverageQuantityBreached
	PreviewFailureReasonPreviewEcosystemLeverageUtilizationBreached
	PreviewFailureReasonPreviewCloseOnlyFailure
	PreviewFailureReasonPreviewUboHighLeverageNotionalBreached
	PreviewFailureReasonPreviewEndTimeTooFarInFuture
	PreviewFailureReasonPreviewLimitPriceTooFarFromMarket
	PreviewFailureReasonPreviewFuturesAfterHourInvalidOrderType
	PreviewFailureReasonPreviewFuturesAfterHourInvalidTimeInForce
	PreviewFailureReasonPreviewInvalidAttachedTakeProfitPrice
	PreviewFailureReasonPreviewInvalidAttachedStopLossPrice
	PreviewFailureReasonPreviewInvalidAttachedTakeProfitPricePrecision
	PreviewFailureReasonPreviewInvalidAttachedStopLossPricePrecision
	PreviewFailureReasonPreviewInvalidAttachedTakeProfitPriceOutOfBounds
	PreviewFailureReasonPreviewInvalidAttachedStopLossPriceOutOfBounds
	PreviewFailureReasonPreviewInvalidBracketOrderSide
	PreviewFailureReasonPreviewBracketOrderSizeExceedsPosition
	PreviewFailureReasonPreviewOrderSizeExceedsBracketedPosition
	PreviewFailureReasonPreviewInvalidLimitPricePrecision
	PreviewFailureReasonPreviewInvalidStopTriggerPricePrecision
	PreviewFailureReasonPreviewInvalidAttachedTakeProfitPriceExceedsMaxDistanceFromOriginatingPrice
	PreviewFailureReasonPreviewInvalidAttachedTakeProfitSizeBelowMin
	PreviewFailureReasonPreviewAttachedOrderSizeMustBeNil
	PreviewFailureReasonPreviewBelowMinSizeForDuration
)

type OrderErrorMetadata struct {
	Error                 *OrderFailureReason   `json:"error"`                    // **(Deprecated)** The reason the order failed to be created
	Message               *string               `json:"message"`                  // Generic error message explaining why the order was not created.
	ErrorDetails          *string               `json:"error_details"`            // Descriptive error message explaining why the order was not created.
	PreviewFailureReason  *PreviewFailureReason `json:"preview_failure_reason"`   // **(Deprecated)** The reason the order failed to be created
	NewOrderFailureReason *OrderFailureReason   `json:"new_order_failure_reason"` // The reason the order failed to be created.
}

// The amount of the second Asset in the Trading Pair.
type QuoteSize struct {
	// The amount of the second Asset in the Trading Pair.
	// For example, on the BTC/USD Order Book, USD is the Quote Asset.
	QuoteSize *string `json:"quote_size,omitempty"`
}

// The amount of the first Asset in the Trading Pair.
type BaseSize struct {
	// The amount of the first Asset in the Trading Pair.
	// For example, on the BTC-USD Order Book, BTC is the Base Asset.
	BaseSize *string `json:"base_size,omitempty"`
}

// Order Good Till Date.
// Order will be canceled if not filled by date.
type GoodTillDate struct {
	// The time at which the order will be cancelled if it is not Filled.
	EndTime *time.Time `json:"end_time,omitempty"`
}

// Embeds stop trigger functionality in an order.
type StopTrigger struct {
	// The price level (in quote currency) where the position will be exited.
	// When triggered, a stop limit order is automatically placed with a limit
	// price 5% higher for BUYS and 5% lower for SELLS.
	StopTriggerPrice *string `json:"stop_trigger_price,omitempty"`
}

type LimitPrice struct {
	// The specified price, or better, that the Order should be executed at.
	// A Buy Order will execute at or lower than the limit price.
	// A Sell Order will execute at or higher than the limit price.
	LimitPrice *string `json:"limit_price,omitempty"`
}

// Enable or disable Post-only Mode.
type PostOnly struct {
	// Enable or disable Post-only Mode. When enabled, only Maker Orders will
	// be posted to the Order Book. Orders that will be posted as a Taker
	// Order will be rejected.
	PostOnly *bool `json:"post_only,omitempty"`
}

type StopOrder struct {
	// The specified price that will trigger the placement of the Order.
	StopPrice *string `json:"stop_price,omitempty"`
	// The direction of the stop limit Order. If Up, then the Order will
	// trigger when the last trade price goes above the stop_price. If Down,
	// then the Order will trigger when the last trade price goes below the
	// stop_price.
	StopDirection *StopDirection `json:"stop_direction,omitempty"`
}

// A time-weighted average price (TWAP) order type that calculates the average
// price of a product to programmatically execute an order over a specified
// duration.
type TimeWeightedAveragePrice struct {
	StartTime      *time.Time     `json:"start_time,omitempty"`      // Time at which the order should begin executing.
	NumberBuckets  *string        `json:"number_buckets,omitempty"`  // The number of smaller buckets/suborders over which the entire order will be broken into. Each suborder will be executed over a duration calculated based on the end_time.
	BucketSize     *string        `json:"bucket_size,omitempty"`     // The size of each suborder. bucket_size multiplied by number_buckets should match the size of the entire twap order).
	BucketDuration *time.Duration `json:"bucket_duration,omitempty"` // The duration over which each sub order was executed.
}

// Market Order Immediate Or Cancel.
// Market orders are used to BUY or SELL a desired product at the given market price. Immediate Or Cancel (ioc): orders instantly cancel the remaining size of the limit order instead of opening it on the book.
type MarketOrderIOC struct {
	BaseSize
	QuoteSize
}

// Buy or sell a specified quantity of an Asset at a specified price.
// The Order will only post to the Order Book if it will immediately Fill;
// any remaining quantity is canceled. Read more on Limit Orders:
// https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#limit-order
type LimitOrderIOC struct {
	BaseSize
	QuoteSize
	LimitPrice
}

// Limit Order Good Till Canceled.
// Limit orders are triggered based on the instructions around quantity and
// price: base_size represents the quantity of your base currency to spend;
// limit_price represents the maximum price at which the order should be filled.
// Good Till Canceled (gtc): orders remain open on the book until canceled.
type LimitOrderGTC struct {
	BaseSize
	QuoteSize
	LimitPrice
	PostOnly
}

// Limit Order Good Till Date.
// Limit orders are triggered based on the instructions around quantity and
// price: base_size represents the quantity of your base currency to spend;
// limit_price represents the maximum price at which the order should be filled.
// Good Till Date (gtd): orders are valid till a specified date or time.
type LimitOrderGTD struct {
	BaseSize
	QuoteSize
	LimitPrice
	PostOnly
	GoodTillDate
}

type TwapLimitGTD struct {
	BaseSize
	QuoteSize
	LimitPrice
	TimeWeightedAveragePrice
	GoodTillDate
}

//go:generate enumer -type=StopDirection -transform=snake-upper -json -text
type StopDirection byte

const (
	StopDirectionStopUp StopDirection = iota
	StopDirectionStopDown
)

// Stop Order Good Till Canceled.
// Stop orders are triggered based on the movement of the last trade price. The last trade price is the last price at which an order was filled.
// Good Till Canceled (gtc): orders remain open on the book until canceled.
type StopLimitOrderGTC struct {
	BaseSize
	LimitPrice
	StopOrder
}

// Stop Order Good Till Date.
// Stop orders are triggered based on the movement of the last trade price. The last trade price is the last price at which an order was filled.
// Good Till Date (gtd): orders are valid till a specified date or time.
type StopLimitOrderGTD struct {
	BaseSize
	LimitPrice
	StopOrder
	GoodTillDate
}

type TriggerBracketGTC struct {
	BaseSize
	LimitPrice
	StopTrigger
}

type TriggerBracketGTD struct {
	BaseSize
	LimitPrice
	StopTrigger
	GoodTillDate
}

// Configuration of the order, it can only consist of a single order type at at time.
// The rest will not be populated.
type OrderConfiguration struct {
	MarketIOC         *MarketOrderIOC    `json:"market_market_ioc,omitempty"`         // Buy or sell a specified quantity of an Asset at the current best available market price. [Read more on Market Orders](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#market-order).
	SorLimitIOC       *LimitOrderIOC     `json:"sor_limit_ioc,omitempty"`             // Buy or sell a specified quantity of an Asset at a specified price. The Order will only post to the Order Book if it will immediately Fill; any remaining quantity is canceled. [Read more on Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#limit-order).
	LimitGTC          *LimitOrderGTC     `json:"limit_limit_gtc,omitempty"`           // Buy or sell a specified quantity of an Asset at a specified price. If posted, the Order will remain on the Order Book until canceled. [Read more on Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#limit-order).
	LimitGTD          *LimitOrderGTD     `json:"limit_limit_gtd,omitempty"`           // Buy or sell a specified quantity of an Asset at a specified price. If posted, the Order will remain on the Order Book until a certain time is reached or the Order is canceled. [Read more on Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#limit-order).
	LimitFOK          *LimitOrderIOC     `json:"limit_limit_fok,omitempty,omitempty"` // Buy or sell a specified quantity of an Asset at a specified price. The Order will only post to the Order Book if it is to immediately and completely Fill. [Read more on Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#limit-order).
	TwapLimitGTD      *TwapLimitGTD      `json:"twap_limit_gtd,omitempty,omitempty"`  // A time-weighted average price (TWAP) order type that calculates the average price of a product to programmatically execute an order over a specified duration.
	StopLimitGTC      *StopLimitOrderGTC `json:"stop_limit_stop_limit_gtc,omitempty"` // Posts an Order to buy or sell a specified quantity of an Asset, but only if and when the last trade price on the Order Book equals or surpasses the Stop Price. If posted, the Order will remain on the Order Book until canceled. [Read more on Stop-Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#stop-limit-order).
	StopLimitGTD      *StopLimitOrderGTD `json:"stop_limit_stop_limit_gtd,omitempty"` // Posts an Order to buy or sell a specified quantity of an Asset, but only if and when the last trade price on the Order Book equals or surpasses the Stop Price. If posted, the Order will remain on the Order Book until a certain time is reached or the Order. [Read more on Stop-Limit Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#stop-limit-order).
	TriggerBracketGTC *TriggerBracketGTC `json:"trigger_bracket_gtc,omitempty"`       // A Limit Order to buy or sell a specified quantity of an Asset at a specified price, with stop limit order parameters embedded in the order. If posted, the Order will remain on the Order Book until canceled. [Read more on Bracket Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#bracket-order).
	TriggerBracketGTD *TriggerBracketGTD `json:"trigger_bracket_gtd,omitempty"`       // A Limit Order to buy or sell a specified quantity of an Asset at a specified price, with stop limit order parameters embedded in the order. If posted, the Order will remain on the Order Book until a certain time is reached or the Order is canceled. [Read more on Bracket Orders.](https://help.coinbase.com/en/coinbase/trading-and-funding/advanced-trade/order-types#bracket-order).
}

// Margin type.
//
//go:generate enumer -type=MarginType -transform=snake-upper -trimprefix=MarginType -json -text
type MarginType byte

const (
	MarginTypeIsolated MarginType = iota
	MarginTypeCross
)

// Edit Order Failure Reason.
//
//go:generate enumer -type=EditFailureReason -transform=snake-upper -trimprefix=EditFailureReason -json -text
type EditFailureReason byte

const (
	EditFailureReasonUnknownOrderEditFailureReason EditFailureReason = iota
	EditFailureReasonCommanderRejectedEditOrder
	EditFailureReasonCannotEditToBelowFilledSize
	EditFailureReasonOrderNotFound
	EditFailureReasonCallerIdMismatch
	EditFailureReasonOnlyLimitOrderEditsSupported
	EditFailureReasonInvalidEditedSize
	EditFailureReasonInvalidEditedPrice
	EditFailureReasonInvalidOriginalSize
	EditFailureReasonInvalidOriginalPrice
	EditFailureReasonEditRequestEqualToOriginalRequest
	EditFailureReasonOnlyOpenOrdersCanBeEdited
	EditFailureReasonSizeInQuoteEditsNotAllowed
	EditFailureReasonOrderIsAlreadyBeingReplaced
)

type EditOrderError struct {
	EditFailureReason    *EditFailureReason    `json:"edit_failure_reason,omitempty"`
	PreviewFailureReason *PreviewFailureReason `json:"preview_failure_reason,omitempty"`
}

// Status of a Coinbase Order.
//
//go:generate enumer -type=OrderStatus -transform=snake-upper -trimprefix=OrderStatus -json -text
type OrderStatus byte

const (
	OrderStatusUnknownOrderStatus OrderStatus = iota
	OrderStatusOpen
	OrderStatusFilled
	OrderStatusCancelled
	OrderStatusExpired
	OrderStatusFailed
	OrderStatusQueued
	OrderStatusCancelQueued
)

// The client specified window for which the order can remain active.
//
//go:generate enumer -type=TimeInForce -transform=snake-upper -trimprefix=TimeInForce -json -text
type TimeInForce byte

const (
	TimeInForceUnknownTimeInForce TimeInForce = iota // Unknown or unspecified.
	TimeInForceGoodUntilDate                         // Orders are valid till a specified date or time.
	TimeInForceGoodUntilCancelled                    //  orders remain open on the book until canceled.
	TimeInForceImmediateOrCancel                     // orders instantly cancel the remaining size of the limit order instead of opening it on the book.
	TimeInForceFillOrKill
)

// The trigger status of an order, with respect to stop price.
//
//go:generate enumer -type=TriggerStatus -transform=snake-upper -trimprefix=TriggerStatus -json -text
type TriggerStatus byte

const (
	TriggerStatusUnknownTriggerStatus TriggerStatus = iota
	TriggerStatusInvalidOrderType
	TriggerStatusStopPending
	TriggerStatusStopTriggered
)

// Type of Coinbase order.
//
//go:generate enumer -type=OrderType -transform=snake-upper -trimprefix=OrderType -json -text
type OrderType byte

const (
	OrderTypeUnknownOrderType OrderType = iota
	OrderTypeMarket
	OrderTypeLimt
	OrderTypeStop
	OrderTypeStopLimit
	OrderTypeBracket
	OrderTypeTwap
)

// Reason an order was rejected.
//
//go:generate enumer -type=OrderRejectReason -transform=snake-upper -trimprefix=OrderRejectReason -json -text
type OrderRejectReason byte

const (
	OrderRejectReasonRejectReasonUnknown OrderRejectReason = iota
	OrderRejectReasonHoldFailure
	OrderRejectReasonTooManyOpenOrders
	OrderRejectReasonRejectReasonInsufficientFunds
	OrderRejectReasonRateLimitExceeded
)

// Source an order was placed from.
//
//go:generate enumer -type=OrderPlacementSource -transform=snake-upper -trimprefix=OrderPlacementSource -json -text
type OrderPlacementSource byte

const (
	OrderPlacementSourceUnknownPlacementSource OrderPlacementSource = iota
	OrderPlacementSourceRetailSimple
	OrderPlacementSourceRetailAdvanced
)

type OrderEditHisory struct {
	Price                  *string    `json:"price,omitempty"` // The update price of the order.
	Size                   *string    `json:"size,omitempty"`  // The updated size of the order.
	ReplaceAcceptTimestamp *time.Time `json:"replace_accept_timestamp,omitempty"`
}

// Coinbase Order.
type Order struct {
	ID                         string                `json:"order_id"`                               // The unique id for this order.
	ProductID                  string                `json:"product_id"`                             // The product this order was created for e.g. 'BTC-USD'.
	UserID                     string                `json:"user_id"`                                // The id of the User owning this Order.
	Configuration              OrderConfiguration    `json:"order_configuration"`                    // Configuration of the order.
	Side                       Side                  `json:"side"`                                   // The side of the market that the order is on (e.g. 'BUY', 'SELL').
	ClientOrderID              string                `json:"client_order_id"`                        // The unique ID provided for the order (used for identification purposes).
	Status                     OrderStatus           `json:"status"`                                 // The current state of the order.
	TimeInForce                *TimeInForce          `json:"time_in_force,omitempty"`                // The client specified window for which the order can remain active.
	CreatedTime                time.Time             `json:"created_time"`                           // Timestamp for when the order was created.
	CompletionPercentage       string                `json:"completion_percentage"`                  // The percent of total order amount that has been filled.
	FilledSize                 *string               `json:"filled_size,omitempty"`                  // The portion (in base currency) of total order amount that has been filled.
	AverageFilledPrice         string                `json:"average_filled_price"`                   // The average of all prices of fills for this order.
	Fee                        *string               `json:"fee,omitempty"`                          // **(Deprecated)** Commission amount.
	NumberOfFills              string                `json:"number_of_fills"`                        // Number of fills that have been posted for this order.
	FilledValue                *string               `json:"filled_value,omitempty"`                 // The amount -- in quote currency -- of the total order that has been filled.
	PendingCancel              bool                  `json:"pending_cancel"`                         // Whether a cancel request has been initiated for the order, and not yet completed.
	SizeInQuote                bool                  `json:"size_in_quote"`                          // Whether the order was placed with quote currency.
	TotalFees                  string                `json:"total_fees"`                             // The total fees for the order.
	SizeInclusiveOfFees        bool                  `json:"size_inclusive_of_fees"`                 // Whether the order size includes fees.
	TotalValueAfterFees        string                `json:"total_value_after_fees"`                 // Derived field defined as (filled_value + total_fees) for buy orders and (filled_value - total_fees) for sell orders.
	TriggerStatus              *TriggerStatus        `json:"trigger_status,omitempty"`               // The trigger status of the order, with respect to stop price.
	Type                       *OrderType            `json:"order_type,omitempty"`                   // Type of the order.
	RejectReason               *OrderRejectReason    `json:"reject_reason,omitempty"`                // Rejection Reason.
	Settled                    *bool                 `json:"settled,omitempty"`                      // True if the order is fully filled, false otherwise.
	ProductType                *ProductType          `json:"product_type,omitempty"`                 // The type of order, i.e. Spot or Future
	RejectMessage              *string               `json:"reject_message,omitempty"`               // Message stating why the order was rejected.
	CancelMessage              *string               `json:"cancel_message,omitempty"`               // Message stating why the order was canceled.
	OrderPlacementSource       *OrderPlacementSource `json:"order_placement_source,omitempty"`       // Message stating which source an order was placed from.
	OutstandingHoldAmount      *string               `json:"outstanding_hold_amount,omitempty"`      // The remaining hold amount calculated as (holdAmount - holdAmountReleased). If the hold is released, returns 0.
	IsLiquidation              *bool                 `json:"is_liquidation,omitempty"`               // True if order is of liquidation type.
	LastFillTime               *time.Time            `json:"last_fill_time,omitempty"`               // Time of the most recent fill for this order.
	EditHistory                []OrderEditHisory     `json:"edit_history,omitempty"`                 // An array of the latest 5 edits per order.
	Leverage                   *string               `json:"leverage,omitempty"`                     // The amount of leverage for the order (default is 1.0).
	MarginType                 *MarginType           `json:"margin_type,omitempty"`                  // Margin Type for this order (default is CROSS).
	RetailPortfolioID          *string               `json:"retail_portfolio_id,omitempty"`          // The ID of the portfolio this order is associated with.
	OriginatingOrderID         *string               `json:"originating_order_id,omitempty"`         // The ID of the parent order of an attached order.
	AttachedORderID            *string               `json:"attached_order_id,omitempty"`            // The ID of the attached order of a parent order.
	AttachedOrderConfiguration *OrderConfiguration   `json:"attached_order_configuration,omitempty"` // The configuration of the attached order. Only TriggerBracketGtc is eligible. Size field must be omitted as the size of the attached order is the same as that of the originating or parent order.
}

// String denoting what type of fill this is.
//
//go:generate enumer -type=FillType -transform=snake-upper -trimprefix=FillType -json -text
type FillType byte

const (
	TradeTypeFill FillType = iota
	TradeTypeReversal
	TradeTypeCorrection
	TradeTypeSynthetic
)

//go:generate enumer -type=LiquidityIndicator -transform=snake-upper -trimprefix=LiquidityIndicator -json -text
type LiquidityIndicator byte

const (
	LiquidityIndicatorUnknownLiquidityIndicator LiquidityIndicator = iota
	LiquidityIndicatorMaker
	LiquidityIndicatorTaker
)

// Fill.
type Fill struct {
	EntryID            *string             `json:"entry_id,omitempty"`            // Unique identifier for the fill.
	TradeID            *string             `json:"trade_id,omitempty"`            // ID of the fill -- unique for all `FILL` trade_types but not unique for adjusted fills.
	OrderID            *string             `json:"order_id,omitempty"`            // ID of the order the fill belongs to.
	TradeTime          *time.Time          `json:"trade_time,omitempty"`          // Time at which this fill was completed.
	TradeType          *FillType           `json:"trade_type,omitempty"`          // String denoting what type of fill this is. Regular fills have the value `FILL`. Adjusted fills have possible values `REVERSAL`, `CORRECTION`, `SYNTHETIC`.
	Price              *string             `json:"price,omitempty"`               // Price the fill was posted at.
	Size               *string             `json:"size,omitempty"`                // Amount of order that was transacted at this fill.
	Commission         *string             `json:"commission,omitempty"`          // Fee amount for fill.
	ProductID          *string             `json:"product_id,omitempty"`          // The trading pair (e.g. 'BTC-USD').
	SequenceTimestamp  *time.Time          `json:"sequence_timestamp,omitempty"`  // Time at which this fill was posted.
	LiquidityIndicator *LiquidityIndicator `json:"liquidity_indicator,omitempty"` // Whether this fill gives or takes liquidity.
	SizeInQuote        *bool               `json:"size_in_quote,omitempty"`       // Whether the order was placed with quote currency.
	UserID             *string             `json:"user_id,omitempty"`             // User that placed the order the fill belongs to.
	Side               *Side               `json:"side,omitempty"`                // Side of order that this fill belongs to.
	RetailPortfolioID  *string             `json:"retail_portfolio_id,omitempty"` // Portfolio that the order fill belongs to.
}

// PNL configuration for a bracket order.
type TriggerBracketPNL struct {
	TakeProfitPNL *string `json:"take_profit_pnl,omitempty"` // PNL if an attached order fills at the take profit price.
	StopLossPNL   *string `json:"stop_loss_pnl,omitempty"`   // PNL if an attached order fills at the stop loss price.
}

// Expected PNL of an order. This value is an estimate and does not take into
// account fees and slippage.
type PnlConfiguration struct {
	TriggerBracketPNL *TriggerBracketPNL `json:"trigger_bracket_pnl,omitempty"` // PNL configuration for a bracket order.
}

// Coinbase Payment method.
type PaymentMethod struct {
	ID            *string    `json:"id,omitempty"`             // Unique identifier for the payment method.
	Type          *string    `json:"type,omitempty"`           // The payment method type.
	Name          *string    `json:"name,omitempty"`           // Name for the payment method.
	Currency      *string    `json:"currency,omitempty"`       // Currency symbol for the payment method.
	Verified      *bool      `json:"verified,omitempty"`       // The verified status of the payment method.
	AllowBuy      *bool      `json:"allow_buy,omitempty"`      // Whether or not this payment method can perform buys.
	AllowSell     *bool      `json:"allow_sell,omitempty"`     // Whether or not this payment method can perform sells.
	AllowDeposit  *bool      `json:"allow_deposit,omitempty"`  // Whether or not this payment method can perform deposits.
	AllowWithdraw *bool      `json:"allow_withdraw,omitempty"` // Whether or not this payment method can perform withdrawals.
	CreatedAt     *time.Time `json:"created_at,omitempty"`     // Time at which this payment method was created.
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`     // Time at which this payment method was updated.
}

// Portfolio margin flag.
//
//go:generate enumer -type=PortfolioMarginFlag -transform=snake-upper -json -text
type PortfolioMarginFlag byte

const (
	PortfolioMarginFlagsUnspecified PortfolioMarginFlag = iota
	PortfolioMarginFlagsInLiquidation
)

// Portfolio liquidation status.
//
//go:generate enumer -type=PortfolioLiquidationStatus -transform=snake-upper -json -text
type PortfolioLiquidationStatus byte

const (
	PortfolioLiquidationStatusUnspecified PortfolioLiquidationStatus = iota
	PortfolioLiquidationStatusCanceling
	PortfolioLiquidationStatusAutoLiquidating
	PortfolioLiquidationStatusLspAssignment
	PortfolioLiquidationStatusCustomerAssignment
	PortfolioLiquidationStatusStatusManual
	PortfolioLiquidationStatusStatusNotLiquidating
)

// Perpetuals portfolio summary.
type PortfolioSummary struct {
	UnrealizedPNL       *AvailableBalance `json:"unrealized_pnl,omitempty"`
	BuyingPower         *AvailableBalance `json:"buying_power,omitempty"`
	TotalBalance        *AvailableBalance `json:"total_balance,omitempty"`
	MaxWithdrawalAmount *AvailableBalance `json:"max_withdrawal_amount,omitempty"`
}

// Perpetuals portfolio.
type PerpetualsPortfolio struct {
	PortfolioUUID              *string                     `json:"portfolio_uuid,omitempty"`           // The portfolio UUID.
	Collateral                 *string                     `json:"collateral,omitempty"`               // The total collateral value in USDC for the portfolio.
	PositionNotional           *string                     `json:"position_notional,omitempty"`        // The total position notional value in USDC for all positions in the portfolio.
	OpenPositionNotional       *string                     `json:"open_position_notional,omitempty"`   // The total position notional value in USDC for all open orders and positions.
	PendingFees                *string                     `json:"pending_fees,omitempty"`             // The accrued fees that has not been paid yet in USDC.
	Borrow                     *string                     `json:"borrow,omitempty"`                   // Total borrow amount in USDC (nets the USDC balance, position PNL, held USDC, accrued interest and rolling debt).
	AccruedInterest            *string                     `json:"accrued_interest,omitempty"`         // Interest charged for borrowed USDC balances.
	RollingDebt                *string                     `json:"rolling_debt,omitempty"`             // Amount of settled transactions that hasn't been paid.
	PortfolioInitialMargin     *string                     `json:"portfolio_initial_margin,omitempty"` // The weighted average of all the position's initial margin utilization.
	PortfolioImNotional        *AvailableBalance           `json:"portfolio_im_notional,omitempty"`
	PortfolioMaintenanceMargin *string                     `json:"portfolio_maintenance_margin,omitempty"` // The maintenance margin of the portfolio.
	PortfolioMmNotional        *AvailableBalance           `json:"portfolio_mm_notional,omitempty"`
	LiquidationPercentage      *string                     `json:"liquidation_percentage,omitempty"` // The liquidation percentage of the portfolio.
	LiquidationBuffer          *string                     `json:"liquidation_buffer,omitempty"`     // The liquidation buffer of the portfolio.
	MarginType                 *MarginType                 `json:"margin_type,omitempty"`
	MarginFlags                *PortfolioMarginFlag        `json:"margin_flags,omitempty"`
	LiquidationStatus          *PortfolioLiquidationStatus `json:"liquidation_status,omitempty"`
	UnrealizedPNL              *AvailableBalance           `json:"unrealized_pnl,omitempty"`
	TotalBalance               *AvailableBalance           `json:"total_balance,omitempty"`
}

// Position Side.
//
//go:generate enumer -type=PositionSide -transform=snake-upper -json -text
type PositionSide byte

const (
	PositionSideUnknown PositionSide = iota
	PositionSideLong
	PositionSideShort
)

type PerpetualsPosition struct {
	ProductID        *string           `json:"product_id,omitempty"`     // The unique identifier of the instrument the position is in.
	ProductUUID      *string           `json:"product_uuid,omitempty"`   // The unique identifier of the instrument the position is in.
	PortfolioUUID    *string           `json:"portfolio_uuid,omitempty"` // The portfolio UUID.
	Symbol           *string           `json:"symbol,omitempty"`         // The trading pair (e.g. 'BTC-PERP-INTX').
	VWAP             *AvailableBalance `json:"vwap,omitempty"`
	EntryVwap        *AvailableBalance `json:"entry_vwap,omitempty"`
	PositionSide     *PositionSide     `json:"position_side,omitempty"`
	MarginType       *MarginType       `json:"margin_type,omitempty"`
	NetSize          *string           `json:"net_size,omitempty"`        // The size of the position with positive values reflecting a long position and negative values reflecting a short position.
	BuyOrderSize     *string           `json:"buy_order_size,omitempty"`  // Cumulative size of all the open buy orders
	SellOrderSize    *string           `json:"sell_order_size,omitempty"` // Cumulative size of all the open sell orders.
	ImContribution   *string           `json:"im_contribution,omitempty"` // The amount this position contributes to the initial margin.
	UnrealizedPNL    *AvailableBalance `json:"unrealized_pnl,omitempty"`
	MarkPrice        *AvailableBalance `json:"mark_price,omitempty"`
	LiquidationPrice *AvailableBalance `json:"liquidation_price,omitempty"`
	Leverage         *string           `json:"leverage,omitempty"` // The leverage of this position.
	ImNotional       *AvailableBalance `json:"im_notional,omitempty"`
	MmNotional       *AvailableBalance `json:"mm_notional,omitempty"`
	PositionNotional *AvailableBalance `json:"position_notional,omitempty"`
	AggregatedPNL    *AvailableBalance `json:"aggregated_pnl,omitempty"`
}

// Perpetuals asset.
type PerpetualsAsset struct {
	AssetID                          *string `json:"asset_id,omitempty"`
	AssetUUID                        *string `json:"asset_uuid,omitempty"`
	AssetName                        *string `json:"asset_name,omitempty"`
	Status                           *string `json:"status,omitempty"`
	CollateralWeight                 *string `json:"collateral_weight,omitempty"`
	AccountCollateralLimit           *string `json:"account_collateral_limit,omitempty"`
	EcosystemCollateralLimitBreached *bool   `json:"ecosystem_collateral_limit_breached,omitempty"`
	AssetIconURL                     *string `json:"asset_icon_url,omitempty"`
	SupportedNetworksAvailable       *bool   `json:"supported_networks_enabled,omitempty"`
}

type PerpetualsBalance struct {
	Asset                        *PerpetualsAsset `json:"asset,omitempty"`
	Quantity                     *string          `json:"quantity,omitempty"`
	Hold                         *string          `json:"hold,omitempty"`
	TransferHold                 *string          `json:"transfer_hold,omitempty"`
	CollateralValue              *string          `json:"collateral_value,omitempty"`
	CollateralWeight             *string          `json:"collateral_weight,omitempty"`
	MaxWithdrawlAmount           *string          `json:"max_withdraw_amount,omitempty"`
	Loan                         *string          `json:"loan,omitempty"`
	LoanCollateralRequirementUSD *string          `json:"loan_collateral_requirement_usd,omitempty"`
	PledgedQuantity              *string          `json:"pledged_quantity,omitempty"`
}

type PerpetualsPortfolioBalance struct {
	PortfolioUUID        *string             `json:"portfolio_uuid,omitempty"`
	Balances             []PerpetualsBalance `json:"balances,omitempty"`
	IsMarginLimitReached bool                `json:"is_margin_limit_reached,omitempty"`
}
