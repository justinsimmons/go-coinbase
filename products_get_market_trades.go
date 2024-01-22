package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type GetMarketTradeOptions struct {
	ProductID string     `url:"-"`                    // The trading pair, i.e., 'BTC-USD'.
	Limit     int        `url:"limit"`                // Number of trades to return.
	Start     *time.Time `url:"start,unix,omitempty"` // Timestamp for starting range of aggregations.
	End       *time.Time `url:"end,unix,omitempty"`   // Timestamp for starting range of aggregations.
}

type Trade struct {
	ID        *string    `json:"trade_id"`   // The ID of the trade that was placed.
	ProductID *string    `json:"product_id"` // The trading pair.
	Price     *string    `json:"price"`      // The price of the trade, in quote currency.
	Size      *string    `json:"size"`       // The size of the trade, in base currency.
	Time      *time.Time `json:"time"`       // The time of the trade.
	Side      *Side      `json:"side"`       // Side of the transaction the trade is on: [BUY, SELL].
	Bid       *string    `json:"bid"`        // The best bid for the `product_id`, in quote currency.
	Ask       *string    `json:"ask"`        // The best ask for the `product_id`, in quote currency.
}

type GetMarketTradesResponse struct {
	Trades  []Trade `json:"trades"`
	BestBid *string `json:"best_bid"` // The best bid for the `product_id`, in quote currency.
	BestAsk *string `json:"best_ask"` // The best ask for the `product_id`, in quote currency.
}

// Get snapshot information, by product ID, about the last trades (ticks), best bid/ask, and 24h volume.
func (s *ProductsService) GetMarketTrades(ctx context.Context, options GetMarketTradeOptions) (*GetMarketTradesResponse, error) {
	url := fmt.Sprintf("%s/api/v3/brokerage/products/%s/ticker", s.client.baseURL, options.ProductID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get market trades HTTP request: %w", err)
	}

	qs, err := query.Values(options)
	if err != nil {
		return nil, fmt.Errorf("failed to convert GetMarketTradeOptions to query string: %w", err)
	}

	req.URL.RawQuery = qs.Encode()

	var trades GetMarketTradesResponse
	err = s.client.do(req, http.StatusOK, &trades)
	if err != nil {
		err = fmt.Errorf("failed to fetch get market trades for product '%s': %w", options.ProductID, err)
	}

	return &trades, err
}
