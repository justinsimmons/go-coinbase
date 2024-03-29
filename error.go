// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"fmt"
	"strings"
)

type ErrorDetails struct {
	// A URL/resource name that uniquely identifies the type of the serialized protocol buffer message.
	// This string must contain at least one "/" character.
	// The last segment of the URL's path must represent the fully qualified name of the type (as in `path/google.protobuf.Duration`).
	// The name should be in a canonical form (e.g., leading "." is not accepted).
	// In practice, teams usually precompile into the binary all types that they expect it to use in the context of Any.
	// However, for URLs which use the scheme `http`, `https`, or no scheme, one can optionally set up a type server that maps type
	// URLs to message definitions as follows: * If no scheme is provided, `https` is assumed. * An HTTP GET on the URL must yield
	// a [google.protobuf.Type][] value in binary format, or produce an error. * Applications are allowed to cache lookup results
	// based on the URL, or have them precompiled into a binary to avoid any lookup. Therefore, binary compatibility needs to be preserved o
	// n changes to types. (Use versioned type names to manage breaking changes.) Note: this functionality is not currently available in
	// the official protobuf release, and it is not used for type URLs beginning with type.googleapis.com. Schemes other than `http`, `https`
	// (or the empty scheme) might be used with implementation specific semantics.
	TypeUrl string `json:"type_url"`
	Value   byte   `json:"value"` // Must be a valid serialized protocol buffer of the above specified type.
}

// Detault error returned by the coinbase API.
type CoinbaseError struct {
	Err     *string        `json:"error"`
	Code    *int32         `json:"code"`
	Message *string        `json:"message"`
	Details []ErrorDetails `json:"details"`
}

func (e CoinbaseError) GetCode() int {
	if e.Code == nil {
		return 0
	}

	return int(*e.Code)
}

func (e CoinbaseError) GetMessage() string {
	if e.Message == nil {
		return ""
	}

	return *e.Message
}

func (e CoinbaseError) getError() string {
	if e.Err == nil {
		return ""
	}

	return *e.Err
}

func (e CoinbaseError) Error() string {
	var details strings.Builder

	details.WriteRune('[') // Always returns nil err.

	n := len(e.Details)
	for i, detail := range e.Details {
		details.WriteString(fmt.Sprintf(`{"type_url": "%s", "value": %v}`, detail.TypeUrl, detail.Value))

		if i < n-1 {
			details.WriteString(", ")
		}
	}

	details.WriteRune(']')

	return fmt.Sprintf(
		`{"error": "%s", "code": %v, "message": "%s", "details": %s}`,
		e.getError(),
		e.GetCode(),
		e.GetMessage(),
		details.String(),
	)
}
