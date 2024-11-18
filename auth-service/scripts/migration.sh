#!/bin/bash
source .env

sleep 2 && goose -dir "migrations" postgres "${MIGRATION_DSN}" up -v