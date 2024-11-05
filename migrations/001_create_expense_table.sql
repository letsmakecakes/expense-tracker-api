CREATE TABLE IF NOT EXISTS expense (
                                     id SERIAL PRIMARY KEY,
                                     date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     description VARCHAR(255) NOT NULL,
                                     amount FLOAT NOT NULL,
                                     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
