// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ErrUnexpectedAPIResponse - coinbase API returned a response outside the API documetation.
var ErrUnexpectedAPIResponse = errors.New("coinbase API returned a response outside the API documetation")

func handleRequestError(resp *http.Response) error {
	if resp == nil {
		return fmt.Errorf("unable to handle null HTTP response")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	var cbError CoinbaseError

	err = json.Unmarshal(body, &cbError)
	if err != nil {
		// Enountered unexpected error from CB.
		errString := ErrUnexpectedAPIResponse.Error()
		status := int32(resp.StatusCode)
		message := string(body)

		return &CoinbaseError{
			Err:     &errString,
			Code:    &status,
			Message: &message,
		}
	}

	return &cbError
}

func (c *Client) do(r *http.Request, successCode int, v any) error {
	// Add required authentication to request.
	c.authenticator.Authenticate(r)

	r.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return fmt.Errorf("failed HTTP request to Coinbase API: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != successCode {
		return handleRequestError(resp)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal HTTP response '%s' into '%T': %w", buf, v, err)
	}

	return nil
}

func (c *Client) get(ctx context.Context, url string, params any, v any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	if params != nil {
		qs, err := query.Values(params)
		if err != nil {
			return fmt.Errorf("failed to convert query params to string: %w", err)
		}

		req.URL.RawQuery = qs.Encode()
	}

	err = c.do(req, http.StatusOK, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) post(ctx context.Context, url string, body io.Reader, v any) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	err = c.do(req, http.StatusOK, v)
	if err != nil {
		return err
	}

	return nil
}
