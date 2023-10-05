-- +goose Up
CREATE TABLE request (
    ID SERIAL PRIMARY KEY,
    URL TEXT NOT NULL,
    Headers JSON NOT NULL,
    Ctx JSON,
    Depth INT NOT NULL,
    Method TEXT NOT NULL,
    Body BYTEA,
    ResponseCharacterEncoding TEXT,
    ProxyURL TEXT
);



-- +goose Down
DROP TABLE request;
