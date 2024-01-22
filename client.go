package coinbase

import "net/http"

const (
	productionURI = "https://api.coinbase.com"
)

// Coinbase Advanced Trade REST API client.
type Client struct {
	apiKey    string // The API key used to authenticate requests (that you create on coinbase.com).
	apiSecret string // The API secret used to authenticate requests (that you create on coinbase.com).

	baseURL    string       // Base URL of the Advanced Trade REST API.
	httpClient *http.Client // Client used to make HTTP calls.

	Accounts *AccountService  // Interface with the Advanced Trade REST API Accounts APIs.
	Orders   *OrdersService   // Interface with the Advanced Trade REST API Orders APIs.
	Products *ProductsService // Interface with the Advanced Trade REST API Bid/Ask API.
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

// New creates a new Coinbase Advanced Trade REST API client.
func New(apiKey string, apiSecret string, opts ...option) *Client {
	c := Client{
		apiKey:    apiKey,
		apiSecret: apiSecret,

		baseURL:    productionURI,
		httpClient: http.DefaultClient,
	}

	// Reuse a single struct instead of allocating one for each service on the heap.
	commonService := service{client: &c}

	c.Accounts = (*AccountService)(&commonService)
	c.Orders = (*OrdersService)(&commonService)
	c.Products = (*ProductsService)(&commonService)

	for _, opt := range opts {
		if opt != nil {
			opt(&c)
		}
	}

	return &c
}
