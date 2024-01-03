-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "secrets"
(
     "id"         BIGSERIAL,
     "user_id"    BIGINT NOT NULL,
     "type_id"    BIGINT NOT NULL,

     "name"       VARCHAR NOT NULL,
     "content"    BYTEA NOT NULL,

     "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
     "updated_at" TIMESTAMP WITH TIME ZONE NULL,
     "deleted_at" TIMESTAMP WITH TIME ZONE NULL
);

ALTER TABLE IF EXISTS "secrets"
    ADD CONSTRAINT "PK_T_SC_C_ID"
        PRIMARY KEY ("id");

ALTER TABLE IF EXISTS "secrets"
    ADD CONSTRAINT "UK_T_SC_C_USER_ID_C_NAME"
        UNIQUE ("user_id", "name");

ALTER TABLE IF EXISTS "secrets"
    ADD CONSTRAINT "FK_T_USR2SC_C_USER_ID"
        FOREIGN KEY ("user_id")
            REFERENCES "users" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "secrets";
-- +goose StatementEnd