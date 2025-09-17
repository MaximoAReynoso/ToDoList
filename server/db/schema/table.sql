CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40) NOT NULL,
    description VARCHAR(255) UNIQUE NOT NULL,
    completed boolean DEFAULT FALSE
)
