-- Drop index
DROP INDEX IF EXISTS idx_users_name;

-- Drop name column
ALTER TABLE users 
DROP COLUMN IF EXISTS name;