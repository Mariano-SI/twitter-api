-- migrate:up
CREATE TABLE IF NOT EXISTS comment_likes (
    id CHAR(36) PRIMARY KEY,
    comment_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_commentlikes_comment FOREIGN KEY (comment_id) REFERENCES comments(id),
    CONSTRAINT fk_commentlikes_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- migrate:down
DROP TABLE IF EXISTS comment_likes;
