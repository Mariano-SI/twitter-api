-- migrate:up
ALTER TABLE post_likes
    ADD CONSTRAINT uq_post_likes_post_user UNIQUE (post_id, user_id);

-- migrate:down
ALTER TABLE post_likes
    DROP INDEX uq_post_likes_post_user;