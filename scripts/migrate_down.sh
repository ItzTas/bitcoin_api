cd coiner_api

if [ -f .env ]; then
    export $(cat .env | xargs)
fi

cd sql/schema

goose postgres "$DB_URL_NO_DISABLE" down