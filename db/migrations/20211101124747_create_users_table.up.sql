CREATE TABLE IF NOT EXISTS users (
    id          uuid DEFAULT uuid_generate_v4(),
    username    VARCHAR(50) NOT NULL,
    created_at  timestamp,
    updated_at  timestamp,
    PRIMARY KEY (id)
);
