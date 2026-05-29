-- migrate:up
CREATE TABLE IF NOT EXISTS post_images (
    id         CHAR(36) PRIMARY KEY,
    post_id    CHAR(36) NOT NULL,
    image_url  VARCHAR(512) NOT NULL UNIQUE,
    position   TINYINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_post_images_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    CONSTRAINT uq_post_images_post_position UNIQUE (post_id, position)
);

-- migrate:down
DROP TABLE IF EXISTS post_images;