-- Drop posts table and indexes
DROP INDEX IF EXISTS idx_posts_user_created_at_id_active;
DROP INDEX IF EXISTS idx_posts_created_at_id_active;
DROP TABLE IF EXISTS posts;
