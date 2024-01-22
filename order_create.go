// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type MarginType string

const (
	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCross    MarginType = "CROSS"
)

type OrderFailureReason string

const (
	OrderFailureReasonUnknown                       OrderFailureReason = "UNKNOWN_FAILURE_REASON"
	OrderFailureReasonUnsupportedOrderConfiguration OrderFailureReason = "UNSUPPORTED_ORDER_CONFIGURATION"
	OrderFailureReasonInvalidSide                   OrderFailureReason = "INVALID_SIDE"
	OrderFailureReasonInvalidProductID              OrderFailureReason = "INVALID_PRODUCT_ID"
	OrderFailureReasonInvalidSizePrecision          OrderFailureReason = "INVALID_SIZE_PRECISION"
	OrderFailureReasonInvalidPricePrecision         OrderFailureReason = "INVALID_PRICE_PRECISION"
	OrderFailureReasonInsufficientFund              OrderFailureReason = "INSUFFICIENT_FUND"
	OrderFailureReasonInvalidLedgerBalance          OrderFailureReason = "INVALID_LEDGER_BALANCE"
	OrderFailureReasonOrderEntryDisabled            OrderFailureReason = "ORDER_ENTRY_DISABLED"
	OrderFailureReasonIneligiblePair                OrderFailureReason = "INELIGIBLE_PAIR"
	OrderFailureReasonInvalidLimitPricePostOnly     OrderFailureReason = "INVALID_LIMIT_PRICE_POST_ONLY"
	OrderFailureReasonInvalidLimitPrice             OrderFailureReason = "INVALID_LIMIT_PRICE"
	OrderFailureReasonInvalidNoLiquidity            OrderFailureReason = "INVALID_NO_LIQUIDITY"
	OrderFailureReasonInvalidRequest                OrderFailureReason = "INVALID_REQUEST"
	OrderFailureReasonCommanderRejectedNewOrder     OrderFailureReason = "COMMANDER_REJECTED_NEW_ORDER"
	OrderFailureReasonInsufficientFunds             OrderFailureReason = "INSUFFICIENT_FUNDS"
)

type CreateOrderOptions struct {
	ClientOrderID         string             `json:"client_order_id"`          // A unique ID provided by the client for their own identification purposes. This ID differs from the order_id generated for the order. If the ID provided is not unique, the order fails to be created and the order corresponding to that ID is returned.
	ProductID             string             `json:"product_id"`               // The product this order was created for e.g. 'BTC-USD'.
	Side                  *Side              `json:"side"`                     // Possible values: [BUY, SELL].
	OrderConfiguration    OrderConfiguration `json:"order_configuration"`      // Configuration of the order details.
	SelfTradePreventionID *string            `json:"self_trade_prevention_id"` // Self trade prevention ID, to prevent an order crossing against the same user.
	Leverage              *string            `json:"leverage"`                 // Leverage for this order; default value is "1.0".
	MarginType            *MarginType        `json:"margin_type"`              // Side the order is on (BUY, SELL).
	RetailPortfolioID     *string            `json:"retail_portfolio_id"`      // Retail portfolio uuid, to associate this order with a retail portfolio.
}

type CreateOrderSuccessMetadata struct {
	OrderID       string  `json:"order_id"`        // The ID of the order created.
	ProductID     *string `json:"product_id"`      // The product this order was created for e.g. 'BTC-USD'.
	Side          *Side   `json:"side"`            // Possible values: [BUY, SELL].
	ClientOrderID *string `json:"client_order_id"` // Client specified ID of order.
}

type CreateOrderErrorMetadata struct {
	Error                 *OrderFailureReason   `json:"error"`
	Message               *string               `json:"message"`       // Generic error message explaining why the order was not created.
	ErrorDetails          *string               `json:"error_details"` // Descriptive error message explaining why the order was not created.
	PreviewFailureReason  *PreviewFailureReason `json:"preview_failure_reason"`
	NewOrderFailureReason *OrderFailureReason   `json:"new_order_failure_reason"`
}

type CreateOrderResponse struct {
	Success            bool                       `json:"success"`        // Whether the order was created.
	OrderFailureReason *OrderFailureReason        `json:"failure_reason"` // Reason for order failure if failure occurs.
	OrderID            *string                    `json:"order_id"`       // The ID of the order created.
	SuccessResponse    CreateOrderSuccessMetadata `json:"success_response"`
	ErrorResponse      CreateOrderErrorMetadata   `json:"error_response"`
	OrderConfiguration *OrderConfiguration        `json:"order_configuration"`
}

func (s *OrdersService) Create(ctx context.Context, order CreateOrderOptions) (any, error) {
	b, err := json.Marshal(&order)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OrderRequest to JSON: %w", err)
	}

	url := s.client.baseURL + "/api/v3/brokerage/orders"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to generate create order HTTP request: %w", err)
	}

	var orderResp CreateOrderResponse
	err = s.client.do(req, http.StatusOK, &orderResp)
	if err != nil {
		err = fmt.Errorf("failed to create order: %w", err)
	}

	return &orderResp, err
}
