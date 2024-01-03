-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS  "secret_types"
(
    "id"    BIGINT,
    "value" VARCHAR NOT NULL
);

ALTER TABLE IF EXISTS "secret_types"
    ADD CONSTRAINT "PK_T_ST_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "secret_types"
    ADD CONSTRAINT "CK_T_ST_C_VALUE"
        CHECK ("value" in ('TEXT', 'BINARY', 'CREDENTIALS', 'CARD'));

ALTER TABLE IF EXISTS "secret_types"
    ADD CONSTRAINT "UK_T_ST_C_VALUE"
        UNIQUE ("value");

INSERT INTO secret_types ("id", "value")
VALUES (1, 'TEXT'),
       (2, 'BINARY'),
       (3, 'CREDENTIALS'),
       (4, 'CARD')
ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "secret_types";
-- +goose StatementEnd