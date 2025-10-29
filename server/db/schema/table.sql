CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40) NOT NULL,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN NOT NULL
)
