# go-coinbase
Go SDK for Coinbase's v3 [Advanced Trade REST API](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-overview)

This does not include the [Advanced Trade WebSocket](https://docs.cloud.coinbase.com/advanced-trade-api/docs/ws-overview).

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
client := coinbase.New("api-key", "api-secret")

// List all accounts for the user.
accounts, err := client.Accounts.List(context.Background(), nil)
```
Some APIs have optional parameters that  can be passed:

```go
client := coinbase.New("api-key", "api-secret")

// Will filter the orders response to only return a specific product.
opt := &ListOrdersOptions{ProductID: coinbase.String("BTC-USD")}

// Gets a list of orders that can be filterd with filtered by optional parameters.
orders, err := client.Orders.List(context.Background(), opt)
```
The services of a client divide the API into logical chunks and correspond to the structure of the Advanced Trade REST API documentation at: https://docs.cloud.coinbase.com/advanced-trade-api/docs/welcome .

NOTE: Using the [context](https://godoc.org/context) package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request. In case there is no context available, then `context.Background()` can be used as a starting point.

## Rate Limits

Advanced Trade API endpoints are throttled by user at 30 requests per second.

Learn more about Coinbase rate limiting at https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-rate-limits .


# Supported

| API | Description | Supported |
| --- | ----------- | --------- |
| [List Accounts](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccounts) | Get a list of authenticated accounts for the current user. | ✅ |
| [Get Account](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccount) |Get a list of information about an account, given an account UUID. | ✅ |
| [Create Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_postorder) | Create an order with a specified product_id (asset-pair), side (buy/sell), etc. | ❌ |
| [Cancel Orders](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_cancelorders) | Initiate cancel requests for one or more orders. | ❌ |
| [Edit Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_editorder) | Edit an order with a specified new `size`, or new `price`. | ❌ |
| [Edit Order Preview](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_previeweditorder) | Simulate an edit order request with a specified new size, or new price, to preview the result of an edit. Only limit order types, with time in force type of good-till-cancelled can be edited. | ❌ |
| [List Orders](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_gethistoricalorders) | Get a list of orders filtered by optional query parameters (`product_id`, `order_status`, etc). | ❌ |
| [List Fills](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfills) | Get a list of fills filtered by optional query parameters (`product_id`, `order_id`, etc). | ❌ |
| [Get Order](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_gethistoricalorder) | Get a single order by order ID. | ❌ |
| [Get Best Bid/Ask](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getbestbidask) | Get the best bid/ask for all products. A subset of all products can be returned instead by using the product_ids input. |  ✅ |
| [Get Product Book](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getbestbidask) | Get a list of bids/asks for a single product. The amount of detail shown can be customized with the limit parameter. | ✅ |
| [List Products](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getproducts) | Get a list of the available currency pairs for trading. | ✅ |
| [Get Product](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getproduct) | Get information on a single product by product ID. | ✅ |
| [Get Product Candles](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getcandles) | Get rates for a single product by product ID, grouped in buckets. | ✅ |
| [Get Market Trades](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getmarkettrades) | Get snapshot information, by product ID, about the last trades (ticks), best bid/ask, and 24h volume. | ✅ |
| [List Portfolios](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getportfolios) | Get a list of all portfolios of a user. | ❌ |
| [Create Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_createportfolio) | Create a portfolio. | ❌ |
| [Move Portfolio Funds](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_moveportfoliofunds) | Transfer funds between portfolios. | ❌ |
| [Get Portfolio Breakdown](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getportfoliobreakdown) | Get the breakdown of a portfolio by portfolio ID. | ❌ |
| [Delete Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_deleteportfolio) | Delete a portfolio by portfolio ID. | ❌ |
| [Edit Portfolio](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_editportfolio) | Modify a portfolio by portfolio ID. | ❌ |
| [Get Futures Balance Summary](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfcmbalancesummary) | Get information on your balances related to Coinbase Financial Markets (CFM) futures trading. | ❌ |
| [List Futures Positions](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfcmpositions) | Get a list of all open positions in CFM futures products. | ❌ |
| [Get Futures Position](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfcmposition) | Get the position of a specific CFM futures product. | ❌ |
| [Schedule Futures Sweep](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_schedulefcmsweep) | Schedule a sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ❌ |
| [List Futures Sweeps](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getfcmsweeps) | Get information on your pending and/or processing requests to sweep funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ❌ |
| [Cancel Pending Futures Sweep](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_cancelfcmsweep) | Cancel your pending sweep of funds from your CFTC-regulated futures account to your Coinbase Inc. USD Spot wallet. | ❌ |
| [Get Transactions Summary](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_gettransactionsummary) | Get a summary of transactions with fee tiers, total volume, and fees. | ✅ |
| [Create Convert Quote](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_createconvertquote) | Create a convert quote with a specified source currency, target currency, and amount. | ❌ |
| [Commit Convert Trade](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_commitconverttrade) | Commits a convert trade with a specified trade ID, source currency, and target currency. | ❌ |
| [Get Convert Trade](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getconverttrade) | Gets a list of information about a convert trade with a specified trade ID, source currency, and target currency. | ❌ |
| [Get Unix Time](https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getunixtime) | Get the current time from the Coinbase Advanced API. | ✅ |

## Order Management

The maximum number of `OPEN` orders allowed per `product_id` is 500. If you have 500 open orders for a `product_id` at submission, new orders placed for that product immediately enter a failed state.

[Advanced API Order Management](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-orders)

## Coinbase Pro

Coinbase Pro has been disabled for use and all customers have been migrated as of December 1, 2023. This was accelerated from a prior announcement of Pro deprecation in 2024.

You cannot use existing Pro API keys to trade with Advanced Trade. See [Migrating from Pro](https://docs.cloud.coinbase.com/advanced-trade-api/docs/migration).

## License

Copyright 2024 Justin Simmons.

This program is released under the [GNU Affero General Public License v3](./LICENSE) or later.
