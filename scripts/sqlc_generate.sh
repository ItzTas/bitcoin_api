#!/bin/bash
cd bitcoin_api/

dotenv -e .env -- sqlc generate