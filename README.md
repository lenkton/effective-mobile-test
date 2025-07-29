# Subscription Service

## Requirements

* docker
* goose (for migrations)

## Getting started

1. Poputlate the `.env` file
   ```sh
   # it is just an example! do not use these values in production!
   cat .env.example > .env
   ```
1. Poputlate the `.env.postgres` file
   ```sh
   # it is just an example! do not use these values in production!
   cat .env.postgres.example > .env.postgres
   ```
1. Start the containers
   ```sh
   docker compose up -d
   ```
1. Run migrations
   ```sh
   goose -dir db/migrations/ postgres \
     "postgresql://app:test@localhost:5432/subscriptions" up
   ```
1. Use the app
   ```sh
   curl localhost:8080/subscriptions
   ```

## API

OpenAPI could be found in
[api/subscription_api.openapi.yml](api/subscription_api.openapi.yml)