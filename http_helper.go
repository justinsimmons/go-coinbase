package coinbase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	// Add reurired authentication to request.
	c.authenticate(r)

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
