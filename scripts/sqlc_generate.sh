#!/bin/bash
cd coiner_api

dotenv -e .env -- sqlc generate