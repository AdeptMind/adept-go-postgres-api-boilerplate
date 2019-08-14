CREATE TABLE IF NOT EXISTS users(
    id serial,
    description VARCHAR (50),
    age integer NOT NULL,
    name VARCHAR (50) NOT NULL,
    email VARCHAR (50) NOT NULL,
    is_admin boolean NOT NULL DEFAULT FALSE,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
