CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS forums (
    id uuid DEFAULT uuid_generate_v4 (),
    name VARCHAR (50) NOT NULL,
    slug VARCHAR (50) UNIQUE NOT NULL,
    description VARCHAR (300),
    PRIMARY KEY (id)
);
