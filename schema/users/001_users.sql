-- +goose Up
CREATE Table Users (
    id UUID PRIMARY KEY 
    created_at TIMESTAMP NOT NULL
    updated_at TIMESTAMP NOT NULL
    username Text NOT NULL
)

-- +goose Down 
DROP Users;