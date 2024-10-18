// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"fmt"
	"net/http"
	"time"
)

const (
	productionURI = "https://api.coinbase.com"
)

// Authenticator adds authentication details to each HTTP request made from the client.
type Authenticator interface {
	Authenticate(*http.Request) error
}

// Coinbase Advanced Trade REST API client.
type Client struct {
	authenticator Authenticator // Handles authentication for the Advanced Trade REST API.

	baseURL    string       // Base URL of the Advanced Trade REST API.
	httpClient *http.Client // Client used to make HTTP calls.

	Accounts       *AccountService        // Interface with the Advanced Trade REST API Accounts APIs.
	Orders         *OrdersService         // Interface with the Advanced Trade REST API Orders APIs.
	Products       *ProductsService       // Interface with the Advanced Trade REST API Products API.
	Fees           *FeesService           // Interface with the Advanced Trade REST API Fees API.
	Portfolio      *PortfoliosService     // Interface with the Advanced Trade REST API Portfolios API.
	Futures        *FuturesService        // Interface with the Advanced Trade REST API Futures API.
	Public         *PublicService         // Interface with the Advanced Trade REST API's Public API.
	Converts       *ConvertsService       // Interface with the Advanced Trade REST API Converts API.
	PaymentMethods *PaymentMethodsService // Interface with the Advanced Trade REST API's Payment Methods API.
}

type service struct {
	client *Client
}

// option is an optional configuration the caller can user to modify the client.
type option func(*Client)

// WithBaseURL overrides the base URL of the Advanced Trade REST API
// on the client.
func WithBaseURL(url string) func(*Client) {
	return func(c *Client) {
		if url == "" {
			return
		}

		c.baseURL = url
	}
}

// WithHTTPClient overrides the default HTTP client used by the client.
func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		if httpClient == nil {
			return
		}

		c.httpClient = httpClient
	}
}

// WithCustomAuthenticator allows the caller to provide custom authentication schema to the client.
// This is useful to hook into the HTTP request and modify it as desired before it is executed.
func WithCustomAuthenticator(authenticator Authenticator) func(*Client) {
	return func(c *Client) {
		c.authenticator = authenticator
	}
}

// NewClient creates a new Coinbase Advanced Trade REST API client.
// By default no authentication schema has been added. Coinbase will block any unauthenticated
// requests.
//   - To use Legacy API key authentication prefer NewWithLegacy()
//   - To use Cloud API Trading Key authentication prefer NewWithCloud()
//   - To use a custom authentication schema use the WithCustomAuthenticator option as an argument to this method.
func NewClient(opts ...option) *Client {
	c := Client{
		baseURL:       productionURI,
		httpClient:    http.DefaultClient,
		authenticator: unauthenticated{}, // Default to unauthenticated user.
	}

	// Reuse a single struct instead of allocating one for each service on the heap.
	commonService := service{client: &c}

	c.Accounts = (*AccountService)(&commonService)
	c.Products = (*ProductsService)(&commonService)
	c.Orders = (*OrdersService)(&commonService)
	c.Portfolio = (*PortfoliosService)(&commonService)
	c.Fees = (*FeesService)(&commonService)
	c.Converts = (*ConvertsService)(&commonService)
	c.Public = (*PublicService)(&commonService)
	c.PaymentMethods = (*PaymentMethodsService)(&commonService)
	c.Futures = (*FuturesService)(&commonService)

	for _, opt := range opts {
		if opt != nil {
			opt(&c)
		}
	}

	return &c
}

// New creates a new Coinbase Advanced Trade REST API client, using legacy API key authentication.
// Please note some of he newer endpoints will not work unless you use the new Cloud API Trading Keys.
func NewWithLegacy(apiKey string, apiSecret string, opts ...option) *Client {
	c := NewClient(opts...)

	c.authenticator = legacyAuthenticator{
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}

	return c
}

// New creates a new Coinbase Advanced Trade REST API client, using Cloud API Trading Keys
// for authentication.
func NewWithCloud(apiKey string, apiSecret string, opts ...option) (*Client, error) {
	c := NewClient(opts...)

	key, err := parsePrivateKey(apiSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid api secret provided: %w", err)
	}

	c.authenticator = cloudAuthenticator{
		apiKey:     apiKey,
		signingKey: key,
	}

	return c, err
}

// Bool is a helper function that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper function that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	return &v
}

// Int64 is a helper function that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 {
	return &v
}

// String is a helper function that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

// Time is a helper function that allocates a new time value
// to store v and returns a pointer to it.
func Time(v time.Time) *time.Time {
	return &v
}
