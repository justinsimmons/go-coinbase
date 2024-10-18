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
)

type CancelOrderFailureReason string

const (
	CancelOrderFailureReasonUnknown          CancelOrderFailureReason = "UNKNOWN_CANCEL_FAILURE_REASON"
	CancelOrderFailureReasonInvalidRequest   CancelOrderFailureReason = "INVALID_CANCEL_REQUEST"
	CancelOrderFailureReasonUnknownOrder     CancelOrderFailureReason = "UNKNOWN_CANCEL_ORDER"
	CancelOrderFailureReasonRejectedOrder    CancelOrderFailureReason = "COMMANDER_REJECTED_CANCEL_ORDER"
	CancelOrderFailureReasonDuplicateRequest CancelOrderFailureReason = "DUPLICATE_CANCEL_REQUEST"
)

type CancelledOrder struct {
	Success       bool                      `json:"success"` // Whether the cancel request was submitted successfully.
	FailureReason *CancelOrderFailureReason `json:"failure_reason"`
	ID            string                    `json:"order_id"` // The IDs of order cancel request was initiated for.
}

type cancelOrdersRequest struct {
	OrderIDs []string `json:"order_ids"`
}

type cancelOrdersResponse struct {
	Results []CancelledOrder `json:"results"` // The result of initiated cancel requests.
}

// Cancel initiates cancel requests for one or more orders.
//
// The maximum number of order_ids that can be cancelled per request is 100.
// This number may be subject to change in emergency, but if a request exceeds the max, then an
// InvalidArgument error code will be returned with an error message denoting the limit
// Too many orderIDs entered, limit is _.
func (s *OrdersService) Cancel(ctx context.Context, ids ...string) ([]CancelledOrder, error) {
	b, err := json.Marshal(&cancelOrdersRequest{OrderIDs: ids})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal cancel order ids to JSON: %w", err)
	}

	u := s.client.baseURL + "/api/v3/brokerage/orders/batch_cancel"

	var cancelResp cancelOrdersResponse
	err = s.client.post(ctx, u, bytes.NewBuffer(b), &cancelResp)
	if err != nil {
		// TODO: Parse and create sentinal error for: "Too many orderIDs entered, limit is _".
		return nil, fmt.Errorf("failed to cancel orders '%v': %w", ids, err)
	}

	return cancelResp.Results, err
}
