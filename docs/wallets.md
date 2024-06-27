# Wallets endpoint

## POST /v1/wallets

This is a authenticated endpoint and needs a JWT token in the header in the format: `Authorization: Bearer {token_string}`

This endpoint will create a wallet with the user id within the jwt token as the owner of the wallet and the sent coin_id as the currency

- Example request

```json
{
  "crypto_type_id": "bitcoin"
}
```

- Example response

```json
{
  "id": "dae1eaf6-2ba8-4ad0-8788-07bc122bfe3a",
  "owner_id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
  "crypto_type_id": "bitcoin",
  "balance_usd": "0",
  "created_at": "2024-06-27T16:44:44.175688Z",
  "updated_at": "2024-06-27T16:44:44.175688Z"
}
```

## PUT /v1/wallets/{coin_id}/coins

This is a authenticated endpoint and needs a JWT token in the header in the format: `Authorization: Bearer {token_string}`

This endpoint will update the user wallet of the given coin id. If the quantity is negative it will take the money from the wallet and put it into the user currency, if it is positive it will deposit the money from the user currency into the wallet

If the user does not have a wallet with the given coin id it will return an error

- Example request

```json
{
  "value": "30"
}
```

- Example response

```json
{
  "id": "b42ee66e-c7b4-491e-83a2-07891882d2b4",
  "wallet_id": "dae1eaf6-2ba8-4ad0-8788-07bc122bfe3a",
  "amount": "30",
  "executed_at": "2024-06-27T16:56:07.997938Z"
}
```
