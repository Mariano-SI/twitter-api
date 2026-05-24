-- migrate:up
ALTER TABLE refresh_tokens ADD COLUMN expires_at TIMESTAMP NOT NULL;

-- migrate:down
ALTER TABLE refresh_tokens DROP COLUMN expires_at;