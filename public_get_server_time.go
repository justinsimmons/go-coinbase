// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

type CoinbaseServerTime struct {
	ISO          *string `json:"iso"`          // An ISO-8601 representation of the timestamp.
	EpochSeconds *string `json:"epochSeconds"` // A second-precision representation of the timestamp.
	EpochMillis  *string `json:"epochMillis"`  // A millisecond-precision representation of the timestamp.
}

// UnixMilli returns the local Time corresponding to the given Unix time, msec milliseconds since January 1, 1970 UTC.
func (ut CoinbaseServerTime) UnixMilli() (time.Time, error) {
	if ut.EpochMillis == nil {
		return time.Time{}, fmt.Errorf("unable to determine time from null EpochMillis")
	}

	i, err := strconv.ParseInt(*ut.EpochMillis, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse epoch-milliseconds: '%s' as int: %w", *ut.EpochMillis, err)
	}

	return time.UnixMilli(i), nil
}

// GetServerTime gets the current time from the Coinbase Advanced API.
// https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getservertime/
func (s *PublicService) GetServerTime(ctx context.Context) (*CoinbaseServerTime, error) {
	var t CoinbaseServerTime

	err := s.client.get(ctx, s.client.baseURL+"/api/v3/brokerage/time", nil, &t)
	if err != nil {
		err = fmt.Errorf("failed to get Coinbase's server time: %w", err)
	}

	return &t, err
}
