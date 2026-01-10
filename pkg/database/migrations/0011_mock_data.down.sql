-- Remove mock data (in reverse order to maintain referential integrity)

-- Delete User Follows
DELETE FROM user_follows;

-- Delete User Players  
DELETE FROM user_players;

-- Delete Users
DELETE FROM users;

-- Delete Ranked Entities
DELETE FROM ranked_entities;

-- Delete Entities
DELETE FROM entity;

-- Delete Catalog Players
DELETE FROM players;

-- Delete Countries
DELETE FROM countries;

-- Delete Managers
DELETE FROM managers;

-- Delete Clubs
DELETE FROM clubs;

-- Delete Positions
DELETE FROM positions;

-- Reset sequences to 1
ALTER SEQUENCE positions_id_seq RESTART WITH 1;
ALTER SEQUENCE clubs_id_seq RESTART WITH 1;
ALTER SEQUENCE managers_id_seq RESTART WITH 1;
ALTER SEQUENCE countries_id_seq RESTART WITH 1;
ALTER SEQUENCE players_id_seq RESTART WITH 1;
ALTER SEQUENCE entity_id_seq RESTART WITH 1;
ALTER SEQUENCE ranked_entities_id_seq RESTART WITH 1;
ALTER SEQUENCE users_id_seq RESTART WITH 1;