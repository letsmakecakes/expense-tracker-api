CREATE TABLE IF NOT EXISTS login (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(50) NOT NULL,
                                     password BYTEA NOT NULL,  -- Use BYTEA to store hashed passwords securely
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Trigger function to auto-update `updated_at` on record modification
CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the function before each update
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON login
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
