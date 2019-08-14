CREATE TABLE IF NOT EXISTS posts(
    id serial,
    user_id integer,
    description VARCHAR (50),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
