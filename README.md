# go-coinbase
Go SDK for Coinbase's v3 [Advanced Trade REST API](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-overview)

This does not include the [Advanced Trade WebSocket](https://docs.cloud.coinbase.com/advanced-trade-api/docs/ws-overview).

## Project Status

Coinbase is pretty rapidly updating and modifing their API's, which makes it hard to keep on top of new features. I am actively using this repo in a few side projects and will maintain it. For now I will not be implementing new endpoints like the converts APIs any time soon, as I do not use them for any of my side projects. I dont really want to invest the time if I am the only one using the SDK. If you would like any new features implemented or encounter a bug please open a Github issue and I will go ahead and jump on them. Thanks!

## Installation

go-coinbase is compatible with Go releases in module mode:

```bash
go get github.com/justinsimmons/go-coinbase
```

## Usage

```go
    import "github.com/justinsimmons/go-coinbase"
```

Alternatively you can use an alias to shorten the invocation.

```go
    import cb "github.com/justinsimmons/go-coinbase"
```

Construct a new Coinbase client, then use the various services on the client to
access different parts of the Advanced Trade REST API. For example:

```go
client := coinbase.NewWithCloud("api-key", "api-secret")

// List all accounts for the user.
accounts, err := client.Accounts.List(context.Background(), nil)
```

Some APIs have optional parameters that  can be passed:

```go
client := coinbase.NewWithCloud("api-key", "api-secret")

// Will filter the orders response to only return a specific product.
opt := &coinbase.ListOrdersOptions{ProductID: coinbase.String("BTC-USD")}

// Gets a list of orders that can be filterd with filtered by optional parameters.
orders, err := client.Orders.List(context.Background(), opt)
```
The services of a client divide the API into logical chunks and correspond to the structure of the Advanced Trade REST API documentation at: https://docs.cloud.coinbase.com/advanced-trade-api/docs/welcome .

NOTE: Using the [context](https://godoc.org/context) package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

# Supported APIs

## Accounts

| API | Description | Supported |
| --- | ----------- | --------- |
| [List Accounts](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccounts) | Get a list of authenticated accounts for the current user. | ✅ |
| [Get Account](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccount) |Get a list of information about an account, given an account UUID. | ✅ |

## Products

| API | Description | Supported |
| --- | ----------- | --------- |
| [Get Best Bid/Ask](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getbestbidask) | Get the best bid/ask for all products. A subset of all products can be returned instead by using the product_ids input. |  ✅ |
| [Get Product Book](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getbestbidask) | Get a list of bids/asks for a single product. The amount of detail shown can be customized with the limit parameter. | ✅ |
| [List Products](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getproducts) | Get a list of the available currency pairs for trading. | ✅ |
| [Get Product](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getproduct) | Get information on a single product by product ID. | ✅ |
| [Get Product Candles](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getcandles) | Get rates for a single product by product ID, grouped in buckets. | ✅ |
| [Get Market Trades](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getmarkettrades) | Get snapshot information, by product ID, about the last trades (ticks), best bid/ask, and 24h volume. | ✅ |

## Orders

| API | Description | Supported |
| --- | ----------- | --------- |
| [Create Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_postorder) | Create an order with a specified product_id (asset-pair), side (buy/sell), etc. | ✅ |
| [Cancel Orders](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_cancelorders) | Initiate cancel requests for one or more orders. | ✅ |
| [Edit Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_editorder) | Edit an order with a specified new `size`, or new `price`. | ✅ |
| [Edit Order Preview](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_previeweditorder) | Simulate an edit order request with a specified new size, or new price, to preview the result of an edit. Only limit order types, with time in force type of good-till-cancelled can be edited. | ✅ |
| [List Orders](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_gethistoricalorders) | Get a list of orders filtered by optional query parameters (`product_id`, `order_status`, etc). | ✅ |
| [List Fills](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfills) | Get a list of fills filtered by optional query parameters (`product_id`, `order_id`, etc). | ✅ |
| [Get Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_gethistoricalorder) | Get a single order by order ID. | ✅ |
| [Preview Order](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_previeworder/) | Preview the results of an order request before sending. | ❌ |

## Portfolios

| API | Description | Supported |
| --- | ----------- | --------- |
| [List Portfolios](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getportfolios) | Get a list of all portfolios of a user. | ✅ |
| [Create Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_createportfolio) | Create a portfolio. | ✅ |
| [Move Portfolio Funds](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_moveportfoliofunds) | Transfer funds between portfolios. | ✅ |
| [Get Portfolio Breakdown](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getportfoliobreakdown) | Get the breakdown of a portfolio by portfolio ID. | ❌ |
| [Delete Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_deleteportfolio) | Delete a portfolio by portfolio ID. | ✅ |
| [Edit Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_editportfolio) | Modify a portfolio by portfolio ID. | ✅ |
| [Close Position](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_closeposition/) | Places an order to close any open positions for a specified product_id. | ❌ |
| [Get Intraday Margin Setting](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getintradaymarginsetting/) | Get the status of whether your account is opted-in to receive increased leverage on futures trades on weekdays from 8am-4pm ET. During these hours, intraday margin rates apply, which are lower than the standard margin requirement needed to hold a futures position overnight. | ❌ |
| [Set Intraday Margin Setting](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_setintradaymarginsetting/) | Setting your margin window to INTRADAY opts you in to receive increased leverage on futures trades on weekdays from 8am-4pm ET. During these hours, intraday margin rates apply, which are lower than the standard margin requirement needed to hold a futures position overnight. You can opt out at anytime. | ❌ |
| [Get Current Margin Window](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getcurrentmarginwindow/) | Get the current margin window to determine whether intraday or overnight margin rates are in effect. | ❌ |
| [Allocate Portfolio](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_allocateportfolio/) | Allocate more funds to an isolated position in your Perpetuals portfolio. | ❌ |
| [Get Perpetuals Portfolio Summary](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getintxportfoliosummary/) | Get a summary of your Perpetuals portfolio | ❌ |
| [List Perpetuals Positions](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getintxpositions/) | Get a list of open positions in your Perpetuals portfolio | ❌ |
| [Get Perpetuals Position](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getintxposition/) | Get a specific open position in your Perpetuals portfolio. | ❌ |
| [Get Portfolio Balances](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getintxbalances/) | Get a list of asset balances on Intx for a given Portfolio. | ❌ |
| [Opt In or Out of Multi Asset Collateral](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_intxmultiassetcollateral/) | Enable or Disable Multi Asset Collateral for a given Portfolio. | ❌ |

## Futures

| API | Description | Supported |
| --- | ----------- | --------- |
| [Get Futures Balance Summary](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmbalancesummary/) | Get information on your balances related to Coinbase Financial Markets (CFM) futures trading. | ✅ |
| [List Futures Positions](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmpositions/) | Get a list of all open positions in CFM futures products. | ✅ |
| [Get Futures Position](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmposition/) | Get the position of a specific CFM futures product | ✅ |
| [Schedule Futures Sweep](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_schedulefcmsweep/) | Schedule a sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ✅ |
| [List Futures Sweeps](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getfcmsweeps/) | Get information on your pending and/or processing requests to sweep funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ✅ |
| [Cancel Pending Futures Sweep](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_cancelfcmsweep/) | Cancel your pending sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ✅ |


## Fees

| API | Description | Supported |
| --- | ----------- | --------- |
| [Get Transactions Summary](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_gettransactionsummary/) | Get a summary of transactions with fee tiers, total volume, and fees. | ✅ |

## Converts 

| API | Description | Supported |
| --- | ----------- | --------- |
| [Create Convert Quote](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_createconvertquote/) | Create a convert quote with a specified source currency, target currency, and amount. | ❌ |
| [Commit Convert Trade](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_commitconverttrade/) | Commits a convert trade with a specified trade ID, source currency, and target currency. | ❌ |
| [Get Convert Trade](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getconverttrade/) | Gets a list of information about a convert trade with a specified trade ID, source currency, and target currency. | ❌ |

## Public 

| API | Description | Supported |
| --- | ----------- | --------- |
| [Get Server Time](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getservertime/) | Get the current time from the Coinbase Advanced API. | ✅ |
| [Get Public Product Book](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproductbook/) | Get a list of bids/asks for a single product. The amount of detail shown can be customized with the limit parameter. | ✅ |
| [List Public Products](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproducts/) | Get a list of the available currency pairs for trading. | ✅ |
| [Get Public Product](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicproduct/) | Get information on a single product by product ID. | ✅ |
| [Get Public Product Candles](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpubliccandles/) | Get rates for a single product by product ID, grouped in buckets. | ✅ |
| [Get Public Market Trades](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpublicmarkettrades/) | Get snapshot information by product ID about the last trades (ticks) and best bid/ask. | ✅ |

## Payment Methods

| API | Description | Supported |
| --- | ----------- | --------- |
| [List Payment Methods](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpaymentmethods/) | Get a list of payment methods for the current user. | ✅ |
| [Get Payment Method](https://docs.cdp.coinbase.com/advanced-trade/reference/retailbrokerageapi_getpaymentmethod/) | Get information about a payment method for the current user. | ✅ |


✅ = Implemented and fully tested.
⚠️ = Implemented but not able to test.
❌ = Not implemented but on roadmap.

## Authentication

Coinbase has multiple authentication schemes available for different APIs. For information on which scheme you should use please consult the [docs](https://docs.cloud.coinbase.com/advanced-trade-api/docs/auth#authentication-schemes).

I would recommend you use [cloud API trading keys](#cloud-api-trading-keys) as they allow for the full functionality of the advance trade APIs.

The following is the compatibility list for the available authentication schemes with `go-coinbase`.

| Scheme | Supported |
| ------ | --------- |
| [Cloud API Trading keys](#cloud-api-trading-keys) | ✅ |
| [OAuth](#oauth) | ⚠️ |
| [Legacy API Keys](#legacy-api-keys) | ✅ |

### Cloud API Trading Keys

Coinbase Cloud supports two API key types, "Trading" keys and "General" keys. **The Advanced API is only compatible with Cloud API Trading keys.**

[Instructions](https://docs.cloud.coinbase.com/advanced-trade-api/docs/auth#creating-trading-keys) on generating a cloud API trading key.

Generate a client with Cloud API Trading credentials:
```go
// Please don't actually hard code these, pass them in via flag.
const (
    apiKey = "organizations/{org_id}/apiKeys/{key_id}"
    apiSecret = "-----BEGIN EC PRIVATE KEY-----\nYOUR PRIVATE KEY\n-----END EC PRIVATE KEY-----\n"
)

// All API requests will use cloud API trading credentials.
client := coinbase.NewWithCloud(apiKey, apiSecret)
```

### Legacy API Keys

Note that the legacy api keys do not suport any of the newer functionality of the Advanced Trade API (Portfolios, etc.).

[Instructions](https://docs.cloud.coinbase.com/advanced-trade-api/docs/auth#creating-legacy-keys) on generating a legacy api key.

Generate a client with legacy credentials:
```go
// All API requests will use legacy API credentials.
client := coinbase.NewWithLegacy("api-key", "api-secret")
```

### OAuth

The OAuth authentication scheme is for applications serving many users, which is outside my use case for this package. To use this authentication scheme you are going to need to [implement a custom authentication scheme](#cusom-authentication-scheme).

### Custom Authentication Scheme

If for whatever reason you would like to use your own authentication method for the API requests you need not rely on the prebuilt functionality provided by this client. You may implem


1. Implement the `coinbase.Authenticator` interface. Example:
    ```go
    type customAuthenticator struct {
        token string
    }

    // customAuthenticator implements coinbase.Authenticator interface.
    func (a customAuthenticator) Authenticate(r *http.Request) error {
        r.Header.Set("Authorization", "Bearer " + a.token)

        return nil
    }
    ```

1. Inject it into the client with the option `coinbase.WithCustomAuthenticator()`.
    ```go
    token := "foo"
    authenticator := customAuthenticator{token: token}

    // Inject custom authenticator into the client.
    client := client.New(coinbase.WithCustomAuthenticator(authenticator))

    // Headers will now have a {"Authentication": {"Bearer foo"}} entry.
    orders, err := client.Orders.List(context.Background(), opt)
    ```

## Rate Limits

Advanced Trade API endpoints are throttled by user at 30 requests per second.

Learn more about Coinbase rate limiting at https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-rate-limits .


## Order Management

The maximum number of `OPEN` orders allowed per `product_id` is 500. If you have 500 open orders for a `product_id` at submission, new orders placed for that product immediately enter a failed state.

[Advanced API Order Management](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-orders)

## Coinbase Pro

Coinbase Pro has been disabled for use and all customers have been migrated as of December 1, 2023. This was accelerated from a prior announcement of Pro deprecation in 2024.

You cannot use existing Pro API keys to trade with Advanced Trade. See [Migrating from Pro](https://docs.cloud.coinbase.com/advanced-trade-api/docs/migration).

## License

Copyright 2024 Justin Simmons.

This program is released under the [GNU Affero General Public License v3](./LICENSE) or later.
