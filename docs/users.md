## Users endpoints

#### GET /v1/users

this endpoint will return a list of all the users in the database
no body params required 

###### query params accepted
- limit: Sets the limit of users returned (defaults to 20)


###### example response
```json

[
    {
        "id": "127490790AOFNJO",
        "user_name": "i_am_amazing",
        "created_at": timeStamp,
        "updated_at": timeStamp,
        "currency": "100000"
    }
]

```

#### "POST /v1/users"

this