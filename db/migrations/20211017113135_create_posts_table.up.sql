CREATE TABLE IF NOT EXISTS posts (
    id          uuid DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) UNIQUE NOT NULL,
    topic_id    uuid NOT NULL,
    author_id   uuid NOT NULL,
    content     TEXT NOT NULL,
    created_at  timestamp,
    updated_at  timestamp,
    PRIMARY KEY (id)
);
