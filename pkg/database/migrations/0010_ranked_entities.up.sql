-- Create ranked_entities table
CREATE TABLE IF NOT EXISTS ranked_entities (
    id BIGSERIAL PRIMARY KEY,
    entity_id BIGINT REFERENCES entity(id) ON DELETE CASCADE,
    player_id BIGINT REFERENCES players(id) ON DELETE CASCADE,
    club_id BIGINT REFERENCES clubs(id) ON DELETE CASCADE,
    manager_id BIGINT REFERENCES managers(id) ON DELETE CASCADE,
    country_id BIGINT REFERENCES countries(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_ranked_entities_entity_id ON ranked_entities(entity_id);
CREATE INDEX IF NOT EXISTS idx_ranked_entities_player_id ON ranked_entities(player_id);
CREATE INDEX IF NOT EXISTS idx_ranked_entities_club_id ON ranked_entities(club_id);
CREATE INDEX IF NOT EXISTS idx_ranked_entities_manager_id ON ranked_entities(manager_id);
CREATE INDEX IF NOT EXISTS idx_ranked_entities_country_id ON ranked_entities(country_id);

-- Create trigger for ranked_entities table
DROP TRIGGER IF EXISTS ranked_entities_set_updated_at ON ranked_entities;
CREATE TRIGGER ranked_entities_set_updated_at
BEFORE UPDATE ON ranked_entities
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
