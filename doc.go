// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

// Coinbase Advanced Trade API SDK for Golang.
//
// Source code available at: https://github.com/justinsimmons/go-coinbase
//
// The Advanced Trade API (or Advanced API) supports programmatic trading and order management with a REST API and WebSocket protocol for real-time market data.
// REST API documentation: https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-overview
// Web Socket protocol: https://docs.cloud.coinbase.com/advanced-trade-api/docs/ws-overview
//
// Coinbase Advanced replaced Coinbase Pro as the "advanced" trading platform.
// The Coinbase Advanced API lets you manage orders, products, and fees with the new v3 endpoints. It does not duplicate core Coinbase functions for account deposits, withdrawals, and transactions. The Coinbase Advanced REST API supports advanced trading features only.
package coinbase
