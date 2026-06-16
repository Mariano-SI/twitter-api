-- migrate:up
CREATE TABLE IF NOT EXISTS follows (
    id         CHAR(36)  PRIMARY KEY,
    follower_id CHAR(36) NOT NULL,
    followed_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_follows_follower FOREIGN KEY (follower_id) REFERENCES users(id),
    CONSTRAINT fk_follows_followed FOREIGN KEY (followed_id) REFERENCES users(id),
    CONSTRAINT uq_follows UNIQUE (follower_id, followed_id)
);

-- migrate:down
DROP TABLE IF EXISTS follows;
