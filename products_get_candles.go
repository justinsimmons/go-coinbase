package coinbase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type TimeGranularity string

const (
	TimeGranularityUnknown        TimeGranularity = "UNKNOWN_GRANULARITY"
	TimeGranularityOneMinute      TimeGranularity = "ONE_MINUTE"
	TimeGranularityFiveMinutes    TimeGranularity = "FIVE_MINUTE"
	TimeGranularityFifteenMinutes TimeGranularity = "FIFTEEN_MINUTE"
	TimeGranularityThirtyMinutes  TimeGranularity = "THIRTY_MINUTE"
	TimeGranularityOneHour        TimeGranularity = "ONE_HOUR"
	TimeGranularityTwoHours       TimeGranularity = "TWO_HOUR"
	TimeGranularitySixHours       TimeGranularity = "SIX_HOUR"
	TimeGranularityOneDay         TimeGranularity = "ONE_DAY"
)

type GetProductCandlesOptions struct {
	ProductID   string          `url:"-"`           // The trading pair.
	Start       time.Time       `url:"start,unix"`  // Timestamp for starting range of aggregations.
	End         time.Time       `url:"end,unix"`    // Timestamp for ending range of aggregations.
	Granularity TimeGranularity `url:"granularity"` // The time slice value for each candle.
}

type Candles struct {
	Start  *string `json:"start"`  // Timestamp for bucket start time, in UNIX time.
	Low    *string `json:"low"`    // Lowest price during the bucket interval.
	High   *string `json:"high"`   // Highest price during the bucket interval.
	Open   *string `json:"open"`   // Opening price (first trade) in the bucket interval.
	Close  *string `json:"close"`  // Closing price (last trade) in the bucket interval.
	Volume *string `json:"volume"` // Volume of trading activity during the bucket interval.
}

type getProductCandlesResponse struct {
	Candles []Candles `json:"candles"`
}

// GetProductCandles gets rates for a single product by product ID, grouped in buckets.
// id: The trading pair.
// start: Timestamp for starting range of aggregations.
// end: Timestamp for ending range of aggregations.
// granularity: The time slice value for each candle.
func (s *ProductsService) GetProductCandles(ctx context.Context, options GetProductCandlesOptions) ([]Candles, error) {
	url := fmt.Sprintf("%s/api/v3/brokerage/products/%s/candles", s.client.baseURL, options.ProductID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get product candles HTTP request: %w", err)
	}

	qs, err := query.Values(options)
	if err != nil {
		return nil, fmt.Errorf("failed to convert GetProductCandlesOptions to query string: %w", err)
	}

	req.URL.RawQuery = qs.Encode()

	var candlesResp getProductCandlesResponse
	err = s.client.do(req, http.StatusOK, &candlesResp)
	if err != nil {
		err = fmt.Errorf("failed to fetch get product candles for product '%s': %w", options.ProductID, err)
	}

	return candlesResp.Candles, err
}