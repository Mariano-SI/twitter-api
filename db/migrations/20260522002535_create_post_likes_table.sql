-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
    id CHAR(36) PRIMARY KEY,
    post_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_postlikes_post FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_postlikes_user FOREIGN KEY (user_id) REFERENCES users(id)
);

-- migrate:down
DROP TABLE IF EXISTS post_likes;

