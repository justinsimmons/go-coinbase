# go-coinbase
Go SDK for Coinbase's Advanced Trade REST API 

This does not include the [Advanced Trade WebSocket]https://docs.cloud.coinbase.com/advanced-trade-api/docs/ws-overview).

[API Documentation](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-overview)

# Supported

| API | Description | Supported |
| --- | ----------- | --------- |
| [List Accounts] (https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccounts) | Get a list of authenticated accounts for the current user. | - |
| [Get Account] (https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_getaccount) |Get a list of information about an account, given an account UUID. | X |
| [Create Order] (https://docs.cloud.coinbase.com/advanced-trade-api/reference/retailbrokerageapi_postorder) | Create an order with a specified product_id (asset-pair), side (buy/sell), etc. | X |

## Order Management

The maximum number of `OPEN` orders allowed per `product_id` is 500. If you have 500 open orders for a `product_id` at submission, new orders placed for that product immediately enter a failed state.

[Advanced API Order Management](https://docs.cloud.coinbase.com/advanced-trade-api/docs/rest-api-orders)

## Rate Limits

Advanced Trade API endpoints are throttled by user at 30 requests per second.

## Coinbase Pro

Coinbase Pro has been disabled for use and all customers have been migrated as of December 1, 2023. This was accelerated from a prior announcement of Pro deprecation in 2024.

You cannot use existing Pro API keys to trade with Advanced Trade. See [Migrating from Pro](https://docs.cloud.coinbase.com/advanced-trade-api/docs/migration).
