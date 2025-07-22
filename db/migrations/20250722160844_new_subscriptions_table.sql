-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subscriptions (
    id INTEGER GENERATED ALWAYS AS IDENTITY,
    service_name VARCHAR(40) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
