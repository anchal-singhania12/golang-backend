-- Create entity table
CREATE TABLE IF NOT EXISTS entity (
    id BIGSERIAL PRIMARY KEY,
    entity_name VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create trigger for entity table
DROP TRIGGER IF EXISTS entity_set_updated_at ON entity;
CREATE TRIGGER entity_set_updated_at
BEFORE UPDATE ON entity
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
