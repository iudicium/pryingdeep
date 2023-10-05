-- name: CreateRequest :one
INSERT INTO request (URL, Headers, Ctx, Depth, Method, Body, ResponseCharacterEncoding, ProxyURL)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING ID;