#!/bin/sh

# ------------------------------------------------------------------------------
# generate env file
# ------------------------------------------------------------------------------
if [ ! -f .env ]; then
    echo 'Creating .env file'
    cp ./env.example ./.env
fi


# ------------------------------------------------------------------------------
# install tool dependencies
# ------------------------------------------------------------------------------
echo 'Installing project tools'
go install -v github.com/go-task/task/v3/cmd/task@latest
go install -tags -v 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install -v github.com/joho/godotenv/cmd/godotenv@latest
go install -v github.com/volatiletech/sqlboiler/v4@latest
go install -v github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

# ------------------------------------------------------------------------------
# install third-party dependencies
# ------------------------------------------------------------------------------
go mod tidy