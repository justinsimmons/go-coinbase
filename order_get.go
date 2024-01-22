package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type OrderStatus string

const (
	OrderStatusOpen      OrderStatus = "OPEN"
	OrderStatusFilled    OrderStatus = "FILLED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
	OrderStatusExpired   OrderStatus = "EXPIRED"
	OrderStatusUnknown   OrderStatus = "UNKNOWN_ORDER_STATUS"
)

type TimeInForce string

const (
	TimeInForceUnknown            TimeInForce = "UNKNOWN_TIME_IN_FORCE" // Unknown or unspecified.
	TimeInForceGoodUntilDate      TimeInForce = "GOOD_UNTIL_DATE_TIME"  // Orders are valid till a specified date or time.
	TimeInForceGoodUntilCancelled TimeInForce = "GOOD_UNTIL_CANCELLED"  //  orders remain open on the book until canceled.
	TimeInForceImmediateOrCancel  TimeInForce = "IMMEDIATE_OR_CANCEL"   // orders instantly cancel the remaining size of the limit order instead of opening it on the book.
)

type TriggerStatus string

const (
	TriggerStatusUnknown          TriggerStatus = "UNKNOWN_TRIGGER_STATUS"
	TriggerStatusInvalidOrderType TriggerStatus = "INVALID_ORDER_TYPE"
	TriggerStatusStopPending      TriggerStatus = "STOP_PENDING"
	TriggerStatusStopTriggered    TriggerStatus = "STOP_TRIGGERED"
)

type OrderType string

const (
	OrderTypeUnknown   OrderType = "UNKNOWN_ORDER_TYPE"
	OrderTypeMarket    OrderType = "MARKET"
	OrderTypeLimt      OrderType = "LIMIT"
	OrderTypeStop      OrderType = "STOP"
	OrderTypeStopLimit OrderType = "STOP_LIMIT"
)

type RejectReason string

const (
	RejectReasonUnknown RejectReason = "REJECT_REASON_UNSPECIFIED"
)

type OrderPlacementSource string

const (
	OrderPlacementSourceRetailSimple   OrderPlacementSource = "RETAIL_SIMPLE"
	OrderPlacementSourceRetailAdvanced OrderPlacementSource = "RETAIL_ADVANCED"
)

type Order struct {
	ID                    string                `json:"order_id"`            // The unique id for this order.
	ProductID             string                `json:"product_id"`          // The product this order was created for e.g. 'BTC-USD'.
	UserID                string                `json:"user_id"`             // The id of the User owning this Order.
	Configuration         *OrderConfiguration   `json:"order_configuration"` // Configuration of the order.
	Side                  *Side                 `json:"side"`                // Side the order is on (BUY, SELL).
	ClientOrderID         string                `json:"client_order_id"`     // Client specified ID of order.
	Status                *OrderStatus          `json:"status"`              // Status of the order.
	TimeInForce           *TimeInForce          `json:"time_in_force"`
	CreatedTime           time.Time             `json:"created_time"`            // Timestamp for when the order was created.
	CompletionPercentage  string                `json:"completion_percentage"`   // The percent of total order amount that has been filled.
	FilledSize            *string               `json:"filled_size"`             // The portion (in base currency) of total order amount that has been filled.
	AverageFilledPrice    string                `json:"average_filled_price"`    // The average of all prices of fills for this order.
	NumberOfFills         string                `json:"number_of_fills"`         // Number of fills that have been posted for this order.
	FilledValue           *string               `json:"filled_value"`            // The amount -- in quote currency -- of the total order that has been filled.
	PendingCancel         bool                  `json:"pending_cancel"`          // Whether a cancel request has been initiated for the order, and not yet completed.
	SizeInQuote           bool                  `json:"size_in_quote"`           // Whether the order was placed with quote currency.
	TotalFees             string                `json:"total_fees"`              // The total fees for the order.
	SizeInclusiveOfFees   bool                  `json:"size_inclusive_of_fees"`  // Whether the order size includes fees.
	TotalValueAfterFees   string                `json:"total_value_after_fees"`  // Derived field defined as (filled_value + total_fees) for buy orders and (filled_value - total_fees) for sell orders.
	TriggerStatus         *TriggerStatus        `json:"trigger_status"`          // Possible values: [UNKNOWN_TRIGGER_STATUS, INVALID_ORDER_TYPE, STOP_PENDING, STOP_TRIGGERED].
	Type                  *OrderType            `json:"order_type"`              // Possible values: [UNKNOWN_ORDER_TYPE, MARKET, LIMIT, STOP, STOP_LIMIT].
	RejectReason          *RejectReason         `json:"reject_reason"`           // Rejection Reason; Possible values: [REJECT_REASON_UNSPECIFIED].
	Settled               *bool                 `json:"settled"`                 // True if the order is fully filled, false otherwise.
	ProductType           *ProductType          `json:"product_type"`            // Possible values: [SPOT, FUTURE].
	RejectMessage         *string               `json:"reject_message"`          // Message stating why the order was rejected.
	CancelMessage         *string               `json:"cancel_message"`          // Message stating why the order was canceled.
	OrderPlacementSource  *OrderPlacementSource `json:"order_placement_source"`  // Possible values: [RETAIL_SIMPLE, RETAIL_ADVANCED].
	OutstandingHoldAmount *string               `json:"outstanding_hold_amount"` // The remaining hold amount calculated as (holdAmount - holdAmountReleased). If the hold is released, returns 0.
	IsLiquidation         *bool                 `json:"is_liquidation"`          // True if order is of liquidation type.
	LastFillTime          *time.Time            `json:"last_fill_time"`          // Time of the most recent fill for this order.
	EditHistory           []struct {
		Price                  *string    `json:"price"`
		Size                   *string    `json:"size"`
		ReplaceAcceptTimestamp *time.Time `json:"replace_accept_timestamp"`
	} `json:"edit_history"` // An array of the latest 5 edits per order.
}

type getOrderResponse struct {
	Order Order `json:"order"`
}

func (s *OrdersService) Get(ctx context.Context, id string) (*Order, error) {
	url := fmt.Sprintf("%s/api/v3/brokerage/orders/historical/%s", s.client.baseURL, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate get account HTTP request: %w", err)
	}

	var orderResp getOrderResponse
	err = s.client.do(req, http.StatusOK, &orderResp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order '%s' for the current user: %w", id, err)
	}

	return &orderResp.Order, err
}
