-- +goose Up
CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    url TEXT
    headers JSON,
     
)