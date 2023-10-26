CREATE TABLE IF NOT EXISTS posts(
    id serial PRIMARY KEY,
    user_id serial,
    text TEXT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);