CREATE TABLE IF NOT EXISTS topics (
    id             uuid DEFAULT uuid_generate_v4(),
    name           VARCHAR(255) NOT NULL,
    slug           VARCHAR(255) UNIQUE NOT NULL,
    forum_id       uuid NOT NULL,
    author_id      uuid NOT NULL,
    last_poster_id uuid NOT NULL,
    posts          INT DEFAULT 0,
    views          INT DEFAULT 0,
    created_at     timestamp,
    updated_at     timestamp,
    PRIMARY KEY (id)
);
