-- Rename badge column back to banner in users table
ALTER TABLE users 
RENAME COLUMN badge TO banner;