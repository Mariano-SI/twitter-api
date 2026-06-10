-- migrate:up
CREATE TABLE IF NOT EXISTS comment_images (
    id         CHAR(36) PRIMARY KEY,
    comment_id    CHAR(36) NOT NULL,
    image_url  VARCHAR(512) NOT NULL UNIQUE,
    position   TINYINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_comment_images_post FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
    CONSTRAINT uq_comment_images_post_position UNIQUE (comment_id, position)
);

-- migrate:down
DROP TABLE IF EXISTS comment_images;