-- Add name column to users table
ALTER TABLE users 
ADD COLUMN IF NOT EXISTS name VARCHAR(255) DEFAULT '';

-- Update existing records to use user_name as name initially (if needed)
UPDATE users SET name = user_name WHERE name = '' OR name IS NULL;

-- Create index for better query performance on name
CREATE INDEX IF NOT EXISTS idx_users_name ON users(name);