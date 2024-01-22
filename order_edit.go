package coinbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type EditFailureReason string

const (
	EditFailureReasonUnknown                      EditFailureReason = "UNKNOWN_EDIT_ORDER_FAILURE_REASON"
	EditFailureReasonRejectedEditOrder            EditFailureReason = "COMMANDER_REJECTED_EDIT_ORDER"
	EditFailureReasonBelowFilledSize              EditFailureReason = "CANNOT_EDIT_TO_BELOW_FILLED_SIZE"
	EditFailureReasonNotFound                     EditFailureReason = "ORDER_NOT_FOUND"
	EditFailureReasonCallerIdMismatch             EditFailureReason = "CALLER_ID_MISMATCH"
	EditFailureReasonOnlyLimitOrderEditsSupported EditFailureReason = "ONLY_LIMIT_ORDER_EDITS_SUPPORTED"
	EditFailureReasonInvalidEditedSize            EditFailureReason = "INVALID_EDITED_SIZE"
	EditFailureReasonInvalidEditedPrice           EditFailureReason = "INVALID_EDITED_PRICE"
	EditFailureReasonInvalidOriginalSize          EditFailureReason = "INVALID_ORIGINAL_SIZE"
	EditFailureReasonInvalidOriginalPrice         EditFailureReason = "INVALID_ORIGINAL_PRICE"
	EditFailureReasonEditEqualToOriginal          EditFailureReason = "EDIT_REQUEST_EQUAL_TO_ORIGINAL_REQUEST"
	EditFailureReasonOnlyOpenOrdersCanBeEdited    EditFailureReason = "ONLY_OPEN_ORDERS_CAN_BE_EDITED"
)

type EditOrderOptions struct {
	OrderID string  `json:"order_id"` // ID of order to edit.
	Price   *string `json:"price"`    // New price for order.
	Size    *string `json:"size"`     // New size for order.
}

type EditOrderResponse struct {
	Success bool `json:"success"` // Whether the order edit request was placed.
	Errors  []struct {
		EditFailureReason    *EditFailureReason    `json:"edit_failure_reason"`
		PreviewFailureReason *PreviewFailureReason `json:"preview_failure_reason"`
	} `json:"errors"` // Details of any errors that may have occured.
}

func (s *OrdersService) edit(ctx context.Context, url string, options EditOrderOptions) (*EditOrderResponse, error) {
	b, err := json.Marshal(&options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal edit order options to JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to generate edit order HTTP request: %w", err)
	}

	var orderResp EditOrderResponse
	err = s.client.do(req, http.StatusOK, &orderResp)
	if err != nil {
		err = fmt.Errorf("failed to edit order: %w", err)
	}

	return &orderResp, err
}

// EditPreview simulates an edit order request with a specified new size, or new price, to preview the result of an edit. Only limit order types, with time in force type of good-till-cancelled can be edited
func (s *OrdersService) EditPreview(ctx context.Context, options EditOrderOptions) (*EditOrderResponse, error) {
	return s.edit(ctx, s.client.baseURL+"/api/v3/brokerage/orders/edit_preview", options)
}

// Edit an order with a specified new size, or new price. Only limit order types, with time in force type of good-till-cancelled can be edited.
//
// Order Priority:
//   - A client can only send an Edit Order request after the previous request for the same order has been fully processed.
//   - CAUTION: You lose your place in line if you increase size or increase/decrease price.
func (s *OrdersService) Edit(ctx context.Context, options EditOrderOptions) (*EditOrderResponse, error) {
	return s.edit(ctx, s.client.baseURL+"/api/v3/brokerage/orders/edit", options)
}
