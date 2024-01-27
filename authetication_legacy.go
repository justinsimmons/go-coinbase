// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	coinbaseAccessKeyHeader       = "CB-ACCESS-KEY"
	coinbaseAccessTimestampHeader = "CB-ACCESS-TIMESTAMP"
	coinbaseAccessSignHeader      = "CB-ACCESS-SIGN"
)

// The CB-ACCESS-SIGN header is generated by creating a sha256 HMAC object using the API secret
// key on the string timestamp + method + requestPath + body.
// 1. Create a signature string by concatenating the values of these query parameters: timestamp + method + requestPath + body
// Concatenate the values of the  following parameters with the + operator: these query parameters: timestamp + method + requestPath + body.
//   - timestamp is the same as the CB-ACCESS-TIMESTAMP header (+/-30 seconds)
//   - method should be UPPER CASE
//   - requestPath is the full path (minus the base URL and query parameters), for example:
//     /api/v3/brokerage/orders/historical/fills
//     /api/v3/brokerage/products/BTC-USD/ticker
//   - body is the request body string -- it is omitted if there is no request body (typically for GET requests)
//
// 2. Create a sha256 HMAC object with your API secret on the signature string.
// 3. Get the hexadecimal string representation of the sha256 HMAC object and pass that in as the CB-ACCESS-SIGN header.
func (c *Client) createSignature(req *http.Request, ts time.Time) (string, error) {
	var buf string

	if req.Body != nil {
		defer req.Body.Close()

		var err error

		buf, err := io.ReadAll(req.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read HTTP request body to add authenitcation headers: %w", err)
		}

		req.Body = io.NopCloser(bytes.NewBuffer(buf))
	}

	s := strconv.FormatInt(ts.Unix(), 10) + strings.ToUpper(req.Method) + req.URL.Path + string(buf)

	h := hmac.New(sha256.New, []byte(c.apiSecret))

	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil)), nil
}

func (c *Client) authenticate(req *http.Request) error {
	if req == nil {
		return nil
	}

	now := time.Now()

	sig, err := c.createSignature(req, now)
	if err != nil {
		return err
	}

	req.Header.Add(coinbaseAccessKeyHeader, c.apiKey)
	req.Header.Add(coinbaseAccessTimestampHeader, strconv.FormatInt(now.Unix(), 10))
	req.Header.Add(coinbaseAccessSignHeader, sig)

	return nil
}