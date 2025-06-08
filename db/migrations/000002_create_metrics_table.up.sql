CREATE TABLE IF NOT EXISTS metrics (
    metric_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    value int,
    player_id INT UNIQUE REFERENCES players(player_id)
);
