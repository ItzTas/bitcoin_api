# Coiner API

This is coiner API! an API that handles the back end of a fictional cryptocurrency trading app

## Endpoints

To see a complete list of the endpoints go to [docs](https://github.com/ItzTas/coiner_api/tree/main/docs)

## :handshake: Contributing

### Clone the repo

```bash
git clone https://github.com/ItzTas/coiner_api
```

### Set the enviroment variables

first get the [coin gecko key](https://www.coingecko.com/pt/developers/painel)

```bash
export PORT={your_port}
export COIN_GECKO_KEY={your_key}
export DELETE_CODE_SECRET={your_secret}
export DB_URL={your_database_url_with_?sslmode=disable}
export DB_URL_NO_DISABLE={your_database_url_}
```

### Build and run the project

```bash
scripts/run_server.sh
```

### Run the tests

```bash
go test ./...
```

### Make a pull request

If you'd like to contribute please open a pull request to the `main` branch
