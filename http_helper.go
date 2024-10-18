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

// doWithAuthentication adds authentication to the HTTP request with the clients configured
// authentication method. An error is returned if a method is not configured. If you wish
// to proceed as an unauthenticated user set the authentication method to unauthenticated{}.
func (c *Client) doWithAuthentication(r *http.Request, successCode int, v any) error {
	// Add required authentication to request.
	if c.authenticator == nil {
		return fmt.Errorf("client is missing authentication method, please regenerate the client with NewClient() to use in an unauthenticated state.")
	}

	c.authenticator.Authenticate(r)

	return c.do(r, successCode, v)
}

func (c *Client) do(r *http.Request, successCode int, v any) error {
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

	err = c.doWithAuthentication(req, http.StatusOK, v)
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

	err = c.doWithAuthentication(req, http.StatusOK, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) put(ctx context.Context, url string, body io.Reader, v any) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	err = c.doWithAuthentication(req, http.StatusOK, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) delete(ctx context.Context, url string) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Typically HTTP DELETE requests do not return a response body.
	// If it does we can catch it in here and then reevaluate.
	v := map[string]any{}

	err = c.doWithAuthentication(req, http.StatusOK, &v)
	if err != nil {
		return err
	}

	return nil
}
