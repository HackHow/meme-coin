CREATE TABLE IF NOT EXISTS meme_coins
(
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(255) UNIQUE NOT NULL,
    description      VARCHAR(255),
    created_at       TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    popularity_score INTEGER                      DEFAULT 0
);
