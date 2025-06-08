CREATE TABLE IF NOT EXISTS projects (
    project_id SERIAL PRIMARY KEY,
    login VARCHAR(255),
    hashed_password VARCHAR(255),
    developer_id INT UNIQUE REFERENCES developers(developer_id)
);
