-- +goose Up
-- +goose StatementBegin
CREATE TABLE test_table (id SERIAL PRIMARY KEY, name TEXT NOT NULL);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE test_table;

-- +goose StatementEnd
