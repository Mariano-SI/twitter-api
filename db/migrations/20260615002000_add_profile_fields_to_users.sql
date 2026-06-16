-- migrate:up
ALTER TABLE users
    ADD COLUMN description         VARCHAR(160)  NULL,
    ADD COLUMN profile_picture_url VARCHAR(512)  NULL,
    ADD COLUMN profile_picture_key VARCHAR(512)  NULL;

-- migrate:down
ALTER TABLE users
    DROP COLUMN description,
    DROP COLUMN profile_picture_url,
    DROP COLUMN profile_picture_key;
