// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import "github.com/google/uuid"

type PortfoliosService service

type PortfolioType string

const (
	PortfolioTypeUndefined PortfolioType = "UNDEFINED"
	PortfolioTypeDefault   PortfolioType = "DEFAULT"
	PortfolioTypeConsumer  PortfolioType = "CONSUMER"
	PortfolioTypeINTX      PortfolioType = "INTX"
)

type FuturesPositionSide string

const (
	FuturesPositionSideUnspecified FuturesPositionSide = "FUTURES_POSITION_SIDE_UNSPECIFIED"
	FuturesPositionSideLong        FuturesPositionSide = "FUTURES_POSITION_SIDE_LONG"
	FuturesPositionSideShort       FuturesPositionSide = "FUTURES_POSITION_SIDE_SHORT"
)

type Portfolio struct {
	Name    *string        `json:"name"`
	UUID    *uuid.UUID     `json:"uuid"`
	Type    *PortfolioType `json:"type"`
	Deleted *bool          `json:"deleted"`
}

type Funds struct {
	Value    string `json:"value"`    // These two values fully represent the monetary amount. Non-localized amount in decimal notation (e.g. "1.234").
	Currency string `json:"currency"` // Currency symbol (USD, BTC, etc). Not an asset UUID.
}
