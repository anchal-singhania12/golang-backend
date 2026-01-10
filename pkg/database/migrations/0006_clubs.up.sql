-- Create clubs table
CREATE TABLE IF NOT EXISTS clubs (
    id BIGSERIAL PRIMARY KEY,
    club_name VARCHAR(255),
    club_image VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create trigger for clubs table
DROP TRIGGER IF EXISTS clubs_set_updated_at ON clubs;
CREATE TRIGGER clubs_set_updated_at
BEFORE UPDATE ON clubs
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
