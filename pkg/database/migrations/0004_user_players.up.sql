-- Create user_players table
CREATE TABLE IF NOT EXISTS user_players (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    player_id BIGINT NOT NULL REFERENCES players(id) ON DELETE CASCADE,
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, player_id)
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_user_players_user_id ON user_players(user_id);
CREATE INDEX IF NOT EXISTS idx_user_players_player_id ON user_players(player_id);

-- Create trigger for user_players table
DROP TRIGGER IF EXISTS user_players_set_updated_at ON user_players;
CREATE TRIGGER user_players_set_updated_at
BEFORE UPDATE ON user_players
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
