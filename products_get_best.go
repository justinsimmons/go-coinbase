package coinbase

import (
	"context"
	"fmt"
	"net/http"
	stdurl "net/url"
)

type GetBestBidAskResponse struct {
	PriceBooks []PriceBook `json:"pricebooks"`
}

// GetBestBidAsk gets the best bid/ask for all products. A subset of all products can be
// returned instead by using the product_ids input.
func (s *ProductsService) GetBest(ctx context.Context, ids ...string) (*GetBestBidAskResponse, error) {
	url := s.client.baseURL + "/api/v3/brokerage/best_bid_ask"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GetBestBidAsk HTTP request: %w", err)
	}

	if len(ids) > 0 {
		req.URL.RawQuery = stdurl.Values{"product_ids": ids}.Encode()
	}

	var bidAsk GetBestBidAskResponse
	err = s.client.do(req, http.StatusOK, &bidAsk)
	if err != nil {
		err = fmt.Errorf("failed to fetch best bid/ask for products '%v': %w", ids, err)
	}

	return &bidAsk, err
}
