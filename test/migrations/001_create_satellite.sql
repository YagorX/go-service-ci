CREATE TABLE IF NOT EXISTS satellite (
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL UNIQUE,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO satellite (name)
VALUES ('moon')
ON CONFLICT (name) DO NOTHING;