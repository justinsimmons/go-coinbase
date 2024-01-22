package coinbase

import (
	"time"
)

type OrdersService service

type StopDirection string

const (
	StopDirectionUp   StopDirection = "STOP_DIRECTION_STOP_UP"
	StopDirectionDown StopDirection = "STOP_DIRECTION_STOP_DOWN"
)

type Side string

const (
	SideBuy  Side = "BUY"
	SideSell Side = "SELL"
)

type PreviewFailureReason string

const (
	PreviewFailureReasonUnknown                        PreviewFailureReason = "UNKNOWN_PREVIEW_FAILURE_REASON"
	PreviewFailureReasonMissingCommissionRate          PreviewFailureReason = "PREVIEW_MISSING_COMMISSION_RATE"
	PreviewFailureReasonInvalidSide                    PreviewFailureReason = "PREVIEW_INVALID_SIDE"
	PreviewFailureReasonInvalidOrderConfig             PreviewFailureReason = "PREVIEW_INVALID_ORDER_CONFIG"
	PreviewFailureReasonInvalidProductID               PreviewFailureReason = "PREVIEW_INVALID_PRODUCT_ID"
	PreviewFailureReasonInvalidSizePrecision           PreviewFailureReason = "PREVIEW_INVALID_SIZE_PRECISION"
	PreviewFailureReasonInvalidPricePrecision          PreviewFailureReason = "PREVIEW_INVALID_PRICE_PRECISION"
	PreviewFailureReasonMissingProductPriceBook        PreviewFailureReason = "PREVIEW_MISSING_PRODUCT_PRICE_BOOK"
	PreviewFailureReasonInvalidLedgerBalance           PreviewFailureReason = "PREVIEW_INVALID_LEDGER_BALANCE"
	PreviewFailureReasonInsufficientLedgerBalance      PreviewFailureReason = "PREVIEW_INSUFFICIENT_LEDGER_BALANCE"
	PreviewFailureReasonInvalidLimitPricePostOnly      PreviewFailureReason = "PREVIEW_INVALID_LIMIT_PRICE_POST_ONLY"
	PreviewFailureReasonInvalidLimitPrice              PreviewFailureReason = "PREVIEW_INVALID_LIMIT_PRICE"
	PreviewFailureReasonInvalidNoLiquidity             PreviewFailureReason = "PREVIEW_INVALID_NO_LIQUIDITY"
	PreviewFailureReasonInsufficientFund               PreviewFailureReason = "PREVIEW_INSUFFICIENT_FUND"
	PreviewFailureReasonInvalidCommissionConfiguration PreviewFailureReason = "PREVIEW_INVALID_COMMISSION_CONFIGURATION"
	PreviewFailureReasonInvalidStopPrice               PreviewFailureReason = "PREVIEW_INVALID_STOP_PRICE"
	PreviewFailureReasonInvalidBaseSizeTooLarge        PreviewFailureReason = "PREVIEW_INVALID_BASE_SIZE_TOO_LARGE"
	PreviewFailureReasonInvalidBaseSizeTooSmall        PreviewFailureReason = "PREVIEW_INVALID_BASE_SIZE_TOO_SMALL"
	PreviewFailureReasonInvalidQuoteSizePrecision      PreviewFailureReason = "PREVIEW_INVALID_QUOTE_SIZE_PRECISION"
	PreviewFailureReasonInvalidQuoteSizeTooLarge       PreviewFailureReason = "PREVIEW_INVALID_QUOTE_SIZE_TOO_LARGE"
	PreviewFailureReasonInvalidPriceTooLarge           PreviewFailureReason = "PREVIEW_INVALID_PRICE_TOO_LARGE"
	PreviewFailureReasonInvalidQuoteSizeTooSmall       PreviewFailureReason = "PREVIEW_INVALID_QUOTE_SIZE_TOO_SMALL"
	PreviewFailureReasonInsufficientFundsForFutures    PreviewFailureReason = "PREVIEW_INSUFFICIENT_FUNDS_FOR_FUTURES"
	PreviewFailureReasonBreachedPriceLimit             PreviewFailureReason = "PREVIEW_BREACHED_PRICE_LIMIT"
	PreviewFailureReasonBreachedAccountPositionLimit   PreviewFailureReason = "PREVIEW_BREACHED_ACCOUNT_POSITION_LIMIT"
	PreviewFailureReasonBreachedCompanyPositionLimit   PreviewFailureReason = "PREVIEW_BREACHED_COMPANY_POSITION_LIMIT"
	PreviewFailureReasonInvalidMarginHealth            PreviewFailureReason = "PREVIEW_INVALID_MARGIN_HEALTH"
	PreviewFailureReasonRiskProxyFailure               PreviewFailureReason = "PREVIEW_RISK_PROXY_FAILURE"
	PreviewFailureReasonUntradableFCMAccountStatus     PreviewFailureReason = "PREVIEW_UNTRADABLE_FCM_ACCOUNT_STATUS"
)

// Market Order Immediate Or Cancel.
// Market orders are used to BUY or SELL a desired product at the given market price. Immediate Or Cancel (ioc): orders instantly cancel the remaining size of the limit order instead of opening it on the book.
type MarketOrderIOC struct {
	QuoteSize *string `json:"quote_size"` // Amount of quote currency to spend on order. Required for BUY orders
	BaseSize  *string `json:"base_size"`  // Amount of base currency to spend on order. Required for SELL orders.
}

// Limit Order Good Till Canceled.
// Limit orders are triggered based on the instructions around quantity and price: base_size represents the quantity of your base currency to spend; limit_price represents the maximum price at which the order should be filled.
// Good Till Canceled (gtc): orders remain open on the book until canceled.
type LimitOrderGTC struct {
	BaseSize   *string `json:"base_size"`   // Amount of base currency to spend on order.
	LimitPrice *string `json:"limit_price"` // Ceiling price for which the order should get filled.
	PostOnly   *bool   `json:"post_only"`   // The post-only flag indicates that the order should only make liquidity. If any part of the order results in taking liquidity, the order will be rejected and no part of it will execute.
}

// Limit Order Good Till Date.
// Limit orders are triggered based on the instructions around quantity and price: base_size represents the quantity of your base currency to spend; limit_price represents the maximum price at which the order should be filled.
// Good Till Date (gtd): orders are valid till a specified date or time.
type LimitOrderGTD struct {
	BaseSize   *string    `json:"base_size"`   // Amount of base currency to spend on order.
	LimitPrice *string    `json:"limit_price"` // Ceiling price for which the order should get filled.
	EndTime    *time.Time `json:"end_time"`    // Time at which the order should be cancelled if it's not filled.
	PostOnly   *bool      `json:"post_only"`   // The post-only flag indicates that the order should only make liquidity. If any part of the order results in taking liquidity, the order will be rejected and no part of it will execute.
}

// Stop Order Good Till Canceled.
// Stop orders are triggered based on the movement of the last trade price. The last trade price is the last price at which an order was filled.
// Good Till Canceled (gtc): orders remain open on the book until canceled.
type StopLimitOrderGTC struct {
	BaseSize      *string        `json:"base_size"`      // Amount of base currency to spend on order.
	LimitPrice    *string        `json:"limit_price"`    // Ceiling price for which the order should get filled.
	StopPrice     *string        `json:"stop_price"`     // Price at which the order should trigger - if stop direction is Up, then the order will trigger when the last trade price goes above this, otherwise order will trigger when last trade price goes below this price.
	StopDirection *StopDirection `json:"stop_direction"` // Possible values: [STOP_DIRECTION_STOP_UP, STOP_DIRECTION_STOP_DOWN].
}

// Stop Order Good Till Date.
// Stop orders are triggered based on the movement of the last trade price. The last trade price is the last price at which an order was filled.
// Good Till Date (gtd): orders are valid till a specified date or time.
type StopLimitOrderGTD struct {
	BaseSize      *string    `json:"base_size"`      // Amount of base currency to spend on order.
	LimitPrice    *string    `json:"limit_price"`    // Ceiling price for which the order should get filled.
	StopPrice     *string    `json:"stop_price"`     // Price at which the order should trigger - if stop direction is Up, then the order will trigger when the last trade price goes above this, otherwise order will trigger when last trade price goes below this price.
	EndTime       *time.Time `json:"end_time"`       // Time at which the order should be cancelled if it's not filled.
	StopDirection *string    `json:"stop_direction"` // Possible values: [STOP_DIRECTION_STOP_UP, STOP_DIRECTION_STOP_DOWN].
}

// Configuration of the order, it can only consist of a single order type at at time.
// The rest will not be populated.
type OrderConfiguration struct {
	MarketIOC    *MarketOrderIOC    `json:"market_market_ioc"`
	LimitGTC     *LimitOrderGTC     `json:"limit_limit_gtc"`
	LimitGTD     *LimitOrderGTD     `json:"limit_limit_gtd"`
	StopLimitGTC *StopLimitOrderGTC `json:"stop_limit_stop_limit_gtc"`
	StopLimitGTD *StopLimitOrderGTD `json:"stop_limit_stop_limit_gtd"`
}
