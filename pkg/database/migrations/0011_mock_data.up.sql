-- Insert mock data for testing APIs

-- Insert Positions
INSERT INTO positions (id, provider_id, position_name, created_at, updated_at) VALUES
(1, 1, 'Forward', NOW(), NOW()),
(2, 2, 'Midfielder', NOW(), NOW()),
(3, 3, 'Defender', NOW(), NOW()),
(4, 4, 'Goalkeeper', NOW(), NOW()),
(5, 5, 'Striker', NOW(), NOW()),
(6, 6, 'Winger', NOW(), NOW());

-- Insert Clubs
INSERT INTO clubs (id, club_name, club_image, created_at, updated_at) VALUES
(1, 'Barcelona FC', 'https://example.com/barcelona-logo.png', NOW(), NOW()),
(2, 'Real Madrid', 'https://example.com/real-madrid-logo.png', NOW(), NOW()),
(3, 'Manchester United', 'https://example.com/man-united-logo.png', NOW(), NOW()),
(4, 'Manchester City', 'https://example.com/man-city-logo.png', NOW(), NOW()),
(5, 'Bayern Munich', 'https://example.com/bayern-logo.png', NOW(), NOW()),
(6, 'Paris Saint-Germain', 'https://example.com/psg-logo.png', NOW(), NOW());

-- Insert Managers
INSERT INTO managers (id, manager_name, manager_image, created_at, updated_at) VALUES
(1, 'Pep Guardiola', 'https://example.com/pep-guardiola.png', NOW(), NOW()),
(2, 'Carlo Ancelotti', 'https://example.com/carlo-ancelotti.png', NOW(), NOW()),
(3, 'Erik ten Hag', 'https://example.com/erik-ten-hag.png', NOW(), NOW()),
(4, 'Thomas Tuchel', 'https://example.com/thomas-tuchel.png', NOW(), NOW()),
(5, 'Xavi Hernandez', 'https://example.com/xavi-hernandez.png', NOW(), NOW()),
(6, 'Luis Enrique', 'https://example.com/luis-enrique.png', NOW(), NOW());

-- Insert Countries
INSERT INTO countries (id, country_name, country_image, created_at, updated_at) VALUES
(1, 'Spain', 'https://example.com/spain-flag.png', NOW(), NOW()),
(2, 'Brazil', 'https://example.com/brazil-flag.png', NOW(), NOW()),
(3, 'Argentina', 'https://example.com/argentina-flag.png', NOW(), NOW()),
(4, 'France', 'https://example.com/france-flag.png', NOW(), NOW()),
(5, 'Germany', 'https://example.com/germany-flag.png', NOW(), NOW()),
(6, 'England', 'https://example.com/england-flag.png', NOW(), NOW()),
(7, 'Portugal', 'https://example.com/portugal-flag.png', NOW(), NOW()),
(8, 'Netherlands', 'https://example.com/netherlands-flag.png', NOW(), NOW());

-- Insert Catalog Players (for catalog management)
INSERT INTO players (id, name, position_id, created_at, updated_at) VALUES
(1, 'Lionel Messi', 1, NOW(), NOW()),
(2, 'Cristiano Ronaldo', 1, NOW(), NOW()),
(3, 'Neymar Jr', 6, NOW(), NOW()),
(4, 'Kylian Mbappé', 1, NOW(), NOW()),
(5, 'Erling Haaland', 5, NOW(), NOW()),
(6, 'Kevin De Bruyne', 2, NOW(), NOW()),
(7, 'Virgil van Dijk', 3, NOW(), NOW()),
(8, 'Thibaut Courtois', 4, NOW(), NOW()),
(9, 'Luka Modrić', 2, NOW(), NOW()),
(10, 'Robert Lewandowski', 5, NOW(), NOW()),
(11, 'Sadio Mané', 6, NOW(), NOW()),
(12, 'Mohamed Salah', 6, NOW(), NOW());

-- Insert Entities
INSERT INTO entity (id, entity_name, created_at, updated_at) VALUES
(1, 'Best Goal Scorer', NOW(), NOW()),
(2, 'Best Playmaker', NOW(), NOW()),
(3, 'Best Defender', NOW(), NOW()),
(4, 'Best Goalkeeper', NOW(), NOW()),
(5, 'Player of the Year', NOW(), NOW()),
(6, 'Young Player of the Year', NOW(), NOW()),
(7, 'Best Manager', NOW(), NOW()),
(8, 'Best Club', NOW(), NOW());

-- Insert Ranked Entities
INSERT INTO ranked_entities (id, entity_id, player_id, club_id, manager_id, country_id, created_at, updated_at) VALUES
(1, 1, 1, 6, NULL, 3, NOW(), NOW()), -- Messi as Best Goal Scorer
(2, 1, 2, NULL, NULL, 7, NOW(), NOW()), -- Ronaldo as Best Goal Scorer
(3, 2, 6, 4, NULL, 2, NOW(), NOW()), -- De Bruyne as Best Playmaker
(4, 3, 7, NULL, NULL, 8, NOW(), NOW()), -- Van Dijk as Best Defender
(5, 4, 8, 2, NULL, 2, NOW(), NOW()), -- Courtois as Best Goalkeeper
(6, 5, 4, 6, NULL, 4, NOW(), NOW()), -- Mbappé as Player of the Year
(7, 7, NULL, NULL, 1, NULL, NOW(), NOW()), -- Pep as Best Manager
(8, 8, NULL, 4, NULL, NULL, NOW(), NOW()); -- Man City as Best Club

-- Insert Users
INSERT INTO users (id, user_name, photo, chant, bio, banner, best_player, best_club, best_manager, best_country, follower_count, following_count, created_at, updated_at) VALUES
(1, 'footballfan123', 'https://example.com/user1-photo.jpg', 'Visca Barça!', 'Barcelona fan since childhood. Love watching Messi play!', 'https://example.com/user1-banner.jpg', 1, 1, 5, 1, 150, 200, NOW(), NOW()),
(2, 'madridista_forever', 'https://example.com/user2-photo.jpg', 'Hala Madrid!', 'Real Madrid is my life. Best club in the world!', 'https://example.com/user2-banner.jpg', 2, 2, 2, 1, 300, 180, NOW(), NOW()),
(3, 'united_red', 'https://example.com/user3-photo.jpg', 'Glory Glory Man United!', 'Manchester United through thick and thin.', 'https://example.com/user3-banner.jpg', 12, 3, 3, 6, 80, 95, NOW(), NOW()),
(4, 'city_blue', 'https://example.com/user4-photo.jpg', 'Come On City!', 'Pep Guardiola is a genius. City till I die!', 'https://example.com/user4-banner.jpg', 6, 4, 1, 6, 120, 110, NOW(), NOW()),
(5, 'psg_ultras', 'https://example.com/user5-photo.jpg', 'Allez Paris!', 'PSG fan from Paris. Mbappé is the future!', 'https://example.com/user5-banner.jpg', 4, 6, 6, 4, 90, 75, NOW(), NOW());

-- Insert User Players (which players each user has selected)
INSERT INTO user_players (user_id, player_id, added_at) VALUES
(1, 1, NOW()), -- footballfan123 has Messi
(1, 3, NOW()), -- footballfan123 has Neymar
(1, 10, NOW()), -- footballfan123 has Lewandowski
(2, 2, NOW()), -- madridista_forever has Ronaldo
(2, 9, NOW()), -- madridista_forever has Modrić
(2, 8, NOW()), -- madridista_forever has Courtois
(3, 12, NOW()), -- united_red has Salah
(3, 7, NOW()), -- united_red has Van Dijk
(4, 6, NOW()), -- city_blue has De Bruyne
(4, 5, NOW()), -- city_blue has Haaland
(5, 4, NOW()), -- psg_ultras has Mbappé
(5, 3, NOW()); -- psg_ultras has Neymar

-- Insert User Follows (who follows whom)
INSERT INTO user_follows (follower_id, following_id) VALUES
(1, 2), -- footballfan123 follows madridista_forever
(1, 4), -- footballfan123 follows city_blue
(2, 1), -- madridista_forever follows footballfan123
(2, 3), -- madridista_forever follows united_red
(3, 4), -- united_red follows city_blue
(3, 5), -- united_red follows psg_ultras
(4, 1), -- city_blue follows footballfan123
(4, 5), -- city_blue follows psg_ultras
(5, 1), -- psg_ultras follows footballfan123
(5, 2); -- psg_ultras follows madridista_forever

-- Update sequences to ensure proper auto-increment values
SELECT setval('positions_id_seq', (SELECT MAX(id) FROM positions));
SELECT setval('clubs_id_seq', (SELECT MAX(id) FROM clubs));
SELECT setval('managers_id_seq', (SELECT MAX(id) FROM managers));
SELECT setval('countries_id_seq', (SELECT MAX(id) FROM countries));
SELECT setval('players_id_seq', (SELECT MAX(id) FROM players));
SELECT setval('entity_id_seq', (SELECT MAX(id) FROM entity));
SELECT setval('ranked_entities_id_seq', (SELECT MAX(id) FROM ranked_entities));
SELECT setval('users_id_seq', (SELECT MAX(id) FROM users));