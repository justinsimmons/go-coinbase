package cb

import (
	"time"

	"github.com/google/uuid"
)

type AccountService service

// Type of coinbase account.
type AccountType string

const (
	AccountTypeUnspecified AccountType = "UNSPECIFIED"
	AccountTypeCrypto      AccountType = "ACCOUNT_TYPE_CRYPTO"
	AccountTypeFiat        AccountType = "ACCOUNT_TYPE_FIAT"
	AccountTypeVault       AccountType = "ACCOUNT_TYPE_VAULT"
)

// Available balance belonging to coinbase account.
type AvailableBalance struct {
	Value    string `json:"value"`    // Amount of currency that this object represents.
	Currency string `json:"currency"` // Denomination of the currency.
}

// TODO: Not 100% sure what this is..?
type Hold struct {
	Value    string `json:"value"`    // Amount of currency that this object represents.
	Currency string `json:"currency"` // Denomination of the currency.
}

// Coinabase account metadata.
type Account struct {
	ID               *uuid.UUID       `json:"uuid"`              // Unique identifier for account.
	Name             *string          `json:"name"`              // Name for the account.
	Currency         *string          `json:"currency"`          // Currency symbol for the account.
	AvailableBalance AvailableBalance `json:"available_balance"` // Available balance of account.
	Default          *bool            `json:"default"`           // Whether or not this account is the user's primary account.
	Active           *bool            `json:"active"`            // Whether or not this account is active and okay to use.
	CreatedAt        *time.Time       `json:"created_at"`        // Time at which this account was created.
	UpdatedAt        *time.Time       `json:"updated_at"`        // Time at which this account was updated.
	DeletedAt        *time.Time       `json:"deleted_at"`        // Time at which this account was deleted.
	Type             *AccountType     `json:"type"`              // Type of account.
	Ready            *bool            `json:"ready"`             // Whether or not this account is ready to trade.
	Hold             Hold             `json:"hold"`
}
