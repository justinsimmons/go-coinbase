// Copyright 2024 Justin Simmons.
//
// This file is part of go-coinbase.
// go-coinbase is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or any later version.
// go-coinbase is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.
// You should have received a copy of the GNU Affero General Public License along with go-coinbase. If not, see <https://www.gnu.org/licenses/>.

package coinbase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Delete deletes a portfolio by portfolio ID.
func (s *PortfoliosService) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.client.delete(ctx, fmt.Sprintf("%s/api/v3/brokerage/portfolios/%s", s.client.baseURL, id.String()))
	if err != nil {
		return fmt.Errorf("failed to delete protfolio '%s': %w", id.String(), err)
	}

	return nil
}
