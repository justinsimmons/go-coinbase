package cb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// ErrUnexpectedAPIResponse - coinbase API returned a response outside the API documetation.
var ErrUnexpectedAPIResponse = errors.New("coinbase API returned a response outside the API documetation")

// Generic unmarshaler for HTTP responses.
func unmarshal(resp *http.Response, in any) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// TODO: Pipe response into string for better logging.

	err = json.Unmarshal(body, in)
	if err != nil {
		return fmt.Errorf("failed to unmarshal HTTP response body into '%T': %w", in, err)
	}

	return nil
}

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

func (c *Client) newRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return req, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Add reurired authentication to request.
	c.authenticate(req)

	req.Header.Add("Accept", "application/json")

	return req, nil
}
