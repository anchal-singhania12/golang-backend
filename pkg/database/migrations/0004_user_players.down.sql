DROP TRIGGER IF EXISTS user_players_set_updated_at ON user_players;
DROP INDEX IF EXISTS idx_user_players_user_id;
DROP INDEX IF EXISTS idx_user_players_player_id;
DROP TABLE IF EXISTS user_players;
