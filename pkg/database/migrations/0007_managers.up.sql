-- Create managers table
CREATE TABLE IF NOT EXISTS managers (
    id BIGSERIAL PRIMARY KEY,
    manager_name VARCHAR(255),
    manager_image VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create trigger for managers table
DROP TRIGGER IF EXISTS managers_set_updated_at ON managers;
CREATE TRIGGER managers_set_updated_at
BEFORE UPDATE ON managers
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
