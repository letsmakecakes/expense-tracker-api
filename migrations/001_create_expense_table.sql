CREATE TABLE IF NOT EXISTS expense (
                                       id SERIAL PRIMARY KEY,
                                       category VARCHAR(50) NOT NULL CHECK (LENGTH(category) > 0), -- Ensure category is not empty
                                       date DATE NOT NULL CHECK (date <= CURRENT_DATE), -- Ensure date is not in the future
                                       description TEXT NOT NULL CHECK (LENGTH(description) > 0), -- Allow longer descriptions while ensuring it's not empty
                                       amount NUMERIC(10, 2) NOT NULL CHECK (amount > 0), -- Ensure the amount is positive
                                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Initial value
                                       UNIQUE (category, date, description, amount) -- Prevent duplicate expense entries
);

-- Create a trigger to automatically update 'updated_at' on row updates
CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON expense
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
