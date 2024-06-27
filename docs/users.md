# Users endpoints

## GET /v1/users

this endpoint will return a list of all the users in the database, add a user id to the end to get a single user /v1/users/{user_id}

no body params required

> Query params

- limit: Sets the limit of users returned (defaults to 20)

> Example response

```json

[
  {
    "id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
    "user_name": "user",
    "created_at": "2024-06-26T22:03:03.426854Z",
    "updated_at": "2024-06-26T22:03:03.426854Z",
    "currency": "0"
  },
  {
    "id": "baa6efc9-e31d-4691-b2a1-379cae3cbb7b",
    "user_name": "user2",
    "created_at": "2024-06-26T22:07:01.467096Z",
    "updated_at": "2024-06-26T22:07:01.467096Z",
    "currency": "0"
  }
]

```

## "POST /v1/users"

this endpoint will create a user to the database and return

> Examples

- Example request

```json
{
  "user_name": "user",
  "email": "exemple@email",
  "password": "password"
}
```

- Example response

```json
{
  "id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
  "user_name": "user",
  "created_at": "2024-06-26T22:03:03.426854Z",
  "updated_at": "2024-06-26T22:03:03.426854Z",
  "currency": "0"
}
```

## POST /v1/login

It will return the user data and a jwt token to use in latter authentications

> Examples

- Example request

```json
{
  "user_name": "user",
  "email": "exemple@email",
  "password": "password"
}
```

- Example response

```json
{
  "id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
  "user_name": "user",
  "created_at": "2024-06-26T22:03:03.426854Z",
  "updated_at": "2024-06-26T22:03:03.426854Z",
  "currency": "100",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiaXRjb2luX2FwaSIsInN1YiI6ImZmNGViNDNmLTljNmItNGRhOC1iMjdjLWNiNzJjNDY3NGI3MyIsImV4cCI6MTcxOTQ2ODg0OSwiaWF0IjoxNzE5NDQwMDQ5fQ.2Dbdk30ijIk7KA0Vz2v4bd0J1l2rN5SATcfKz2u5TvI"
}
```

## PUT /v1/users

This is a authenticated endpoint and needs a JWT token in the header in the format: `Authorization: Bearer {token_string}`

This endpoint will update the user password, email and username. To keep one of the fields as it is just omit the camp

- Example request

```json
{
  "new_user_name": "user",
  "new_email": "exemple@email",
  "new_password": "password"
}
```

- responds with 204 code

## POST /v1/users/{receiver_id}/transactions

This is a authenticated endpoint and needs a JWT token in the header in the format: `Authorization: Bearer {token_string}`

This endpoint will send money from one user to another, update their currency and create a transaction in the database

> Examples

- Example request

```json
{
  "send_quantity": "10"
}
```

- Example response

```json
{
  "id": "ba0a67f5-c25c-472c-8874-c5cda38d034f",
  "sender_id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
  "receiver_id": "baa6efc9-e31d-4691-b2a1-379cae3cbb7b",
  "amount": "10",
  "executed_at": "2024-06-26T22:31:32.094246Z"
}
```

## GET /v1/users/{user_id}/transactions

This endpoint returns every transaction from one user

> Query params

- Limit param

- Example response

```json
[
  {
    "id": "ba0a67f5-c25c-472c-8874-c5cda38d034f",
    "sender_id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
    "receiver_id": "baa6efc9-e31d-4691-b2a1-379cae3cbb7b",
    "amount": "10",
    "executed_at": "2024-06-26T22:31:32.094246Z",
    "user_role": "sender"
  },
  {
    "id": "2e878ac9-b7b8-479e-bffd-9146f690d417",
    "sender_id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
    "receiver_id": "baa6efc9-e31d-4691-b2a1-379cae3cbb7b",
    "amount": "20",
    "executed_at": "2024-06-27T16:37:14.034157Z",
    "user_role": "sender"
  },
  {
    "id": "3437d6a6-46d4-47fc-b4e8-f64ff09ce0c2",
    "sender_id": "baa6efc9-e31d-4691-b2a1-379cae3cbb7b",
    "receiver_id": "ff4eb43f-9c6b-4da8-b27c-cb72c4674b73",
    "amount": "66",
    "executed_at": "2024-06-27T16:40:49.036291Z",
    "user_role": "receiver"
  }
]
```
