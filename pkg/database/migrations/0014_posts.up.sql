-- Create posts table
CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    images JSONB NOT NULL DEFAULT '[]'::jsonb,
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    is_blocked BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT fk_posts_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for performance
-- Index for fetching posts sorted by created_at and id (for main feed)
CREATE INDEX IF NOT EXISTS idx_posts_created_at_id_active 
ON posts (created_at DESC, id DESC) 
WHERE is_deleted = false AND is_blocked = false;

-- Index for fetching posts by user_id sorted by created_at and id
CREATE INDEX IF NOT EXISTS idx_posts_user_created_at_id_active 
ON posts (user_id, created_at DESC, id DESC) 
WHERE is_deleted = false AND is_blocked = false;
