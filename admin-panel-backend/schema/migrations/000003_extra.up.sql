CREATE TABLE extra_messages (
    id serial not null unique,
    title varchar(255) not null,
    description varchar(512) not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);