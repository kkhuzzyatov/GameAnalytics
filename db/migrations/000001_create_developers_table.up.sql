CREATE TABLE IF NOT EXISTS developers (
    developer_id SERIAL PRIMARY KEY,
    login VARCHAR(255),
    hashed_password VARCHAR(255)
);
