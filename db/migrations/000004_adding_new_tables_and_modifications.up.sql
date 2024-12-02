CREATE TABLE IF NOT EXISTS PARAMETERS(
    PARAMETERS_ID SERIAL PRIMARY KEY,
    MIN_COUNT INTEGER NOT NULL,
    MAX_COUNT INTEGER NOT NULL,
    DEADLINE DATE
);

ALTER TABLE IF EXISTS TEAMS ADD COLUMN IF NOT EXISTS PARAMETERS_ID INTEGER;

ALTER TABLE IF EXISTS TEAMS ADD FOREIGN KEY (PARAMETERS_ID) REFERENCES PARAMETERS