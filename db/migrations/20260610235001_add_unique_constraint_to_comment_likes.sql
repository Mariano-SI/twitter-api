-- migrate:up
ALTER TABLE comment_likes
    ADD CONSTRAINT uq_comment_likes_post_user UNIQUE (comment_id, user_id);

-- migrate:down
ALTER TABLE comment_likes
    DROP INDEX uq_comment_likes_post_user;