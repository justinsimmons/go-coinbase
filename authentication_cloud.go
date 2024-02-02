// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	issuer         = "coinbase-cloud"        // Issuer of the service.
	serviceName    = "retail_rest_api_proxy" // Name of the Coinbase service we are generating the token for.
	jwtTTL         = time.Minute * 2         // Coinbase specifies JWTs will expire after two minutes, after which all requests are unauthenticated.
	creationBuffer = time.Second * -45       // Minimum buffer required for coinbase to accept an auth tokens issued time.
)

// Handles cloud API key authentication used to access the Advanced Trade API.
type cloudAuthenticator struct {
	apiKey     string        // The API key used to authenticate requests.
	signingKey crypto.Signer // Signing key, used to sign JWTs if using cloud trading keys.
}

func parsePrivateKey(secret string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(secret))
	if block == nil {
		return nil, fmt.Errorf("failed to decode private key")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return key, nil
}

// Authenticate adds required cloud API authentication headers to the HTTP request.
func (a cloudAuthenticator) Authenticate(req *http.Request) error {
	nonce, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return fmt.Errorf("failed to generate JWT nonce: %w", err)
	}

	// If we use the time right now coinbase will reject the token.
	// Need to add a negative buffer to the time in order for it to be accepted.
	// Super annoying....
	now := time.Now().Add(creationBuffer)

	t := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"sub": a.apiKey,
			"iss": issuer,
			// "iat": now.Unix(),
			"nbf": now.Unix(),
			"exp": now.Add(jwtTTL).Unix(),
			"aud": serviceName,
			"uri": fmt.Sprintf("%s %s%s", req.Method, req.URL.Host, req.URL.Path),
		},
	)

	t.Header["kid"] = a.apiKey
	t.Header["nonce"] = nonce.String()

	s, err := t.SignedString(a.signingKey)
	if err != nil {
		return fmt.Errorf("failed to sign JWT token: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+s)

	return nil
}
