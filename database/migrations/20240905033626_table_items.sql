-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    description VARCHAR(30) NOT NULL,
    quantity INT,
    order_id INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE OR REPLACE FUNCTION setTimeStamp_Items()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp_items
BEFORE UPDATE ON items
FOR EACH ROW
EXECUTE FUNCTION setTimeStamp_Items();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp_items ON items;
DROP FUNCTION IF EXISTS setTimeStamp_Items();
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
