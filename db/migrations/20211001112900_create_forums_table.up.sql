CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS forums (
    id          uuid DEFAULT uuid_generate_v4 (),
    name        VARCHAR (255) NOT NULL,
    slug        VARCHAR (255) UNIQUE NOT NULL,
    description VARCHAR (255),
    author_id   uuid NOT NULL,
    topics      INT DEFAULT 0,
    posts       INT DEFAULT 0,
    parent_id   uuid DEFAULT NULL,
    created_at  timestamp,
    updated_at  timestamp,
    PRIMARY KEY (id)
);
