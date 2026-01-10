# FanligaFC API Testing Guide

## 🚀 Getting Started

The database has been populated with comprehensive mock data to test all APIs.

**Start the server:**

```bash
CONFIG_PATH=config/local.yaml go run cmd/main.go
```

Server runs on: `http://localhost:3000`

---

## 📊 Mock Data Summary

### Users (5 users)

| ID | Username | Best Player | Best Club | Best Manager | Best Country |
|----|----------|-------------|-----------|--------------|--------------|
| 1 | footballfan123 | Messi (1) | Barcelona (1) | Xavi (5) | Spain (1) |
| 2 | madridista_forever | Ronaldo (2) | Real Madrid (2) | Ancelotti (2) | Spain (1) |
| 3 | united_red | Salah (12) | Man United (3) | ten Hag (3) | England (6) |
| 4 | city_blue | De Bruyne (6) | Man City (4) | Guardiola (1) | England (6) |
| 5 | psg_ultras | Mbappé (4) | PSG (6) | Luis Enrique (6) | France (4) |

### Players (12 players)

| ID | Name | Position |
|----|------|----------|
| 1 | Lionel Messi | Forward |
| 2 | Cristiano Ronaldo | Forward |
| 3 | Neymar Jr | Winger |
| 4 | Kylian Mbappé | Forward |
| 5 | Erling Haaland | Striker |
| 6 | Kevin De Bruyne | Midfielder |
| 12 | Mohamed Salah | Winger |

### Clubs (6 clubs)

- Barcelona FC, Real Madrid, Manchester United, Manchester City, Bayern Munich, PSG

### Managers (6 managers)

- Pep Guardiola, Carlo Ancelotti, Erik ten Hag, Thomas Tuchel, Xavi Hernandez, Luis Enrique

### Countries (8 countries)

- Spain, Brazil, Argentina, France, Germany, England, Portugal, Netherlands

---

## 🧪 API Testing Examples

### 1. Health Check

```bash
GET http://localhost:3000/ping
```

### 2. User APIs

#### Fetch User Profile (with joins!)

```bash
GET http://localhost:3000/user/profile
Headers:
  user-id: 1
```

**Expected Response:** Complete user profile with nested best_player, best_club, best_manager, and
best_country data!

#### Fetch User Players

```bash
GET http://localhost:3000/user/players  
Headers:
  user-id: 1
```

**Expected:** User 1 has Messi, Neymar, and Lewandowski

#### Create New User

```bash
POST http://localhost:3000/user/create
Content-Type: application/json

{
  "user_name": "newuser123",
  "photo": "https://example.com/photo.jpg",
  "chant": "Go Team!",
  "bio": "Football enthusiast",
  "banner": "https://example.com/banner.jpg",
  "best_player_id": 1,
  "best_club_id": 1,
  "best_manager_id": 1,
  "best_country_id": 1
}
```

#### Add Players to User

```bash
POST http://localhost:3000/user/players
Headers:
  user-id: 1
  Content-Type: application/json

{
  "player_ids": [7, 8]
}
```

#### Follow/Unfollow User

```bash
POST http://localhost:3000/user/follow
Headers:
  user-id: 1
  Content-Type: application/json

{
  "following_id": 2
}
```

### 3. Catalog APIs

#### List All Positions

```bash
GET http://localhost:3000/catalog/positions
```

#### Get Specific Club

```bash
GET http://localhost:3000/catalog/clubs/1
```

**Expected:** Barcelona FC details

#### Create New Manager

```bash
POST http://localhost:3000/catalog/managers
Content-Type: application/json

{
  "manager_name": "Jurgen Klopp",
  "manager_image": "https://example.com/klopp.jpg"
}
```

#### Update Country

```bash
PUT http://localhost:3000/catalog/countries/1
Content-Type: application/json

{
  "country_name": "Kingdom of Spain",
  "country_image": "https://example.com/spain-updated.png"
}
```

#### Delete Entity

```bash
DELETE http://localhost:3000/catalog/entities/1
```

### 4. Player Management APIs

#### Create Player (with provider_player_id)

```bash
POST http://localhost:3000/player/
Content-Type: application/json

{
  "name": "Vinicius Junior",
  "position_id": 6,
  "provider_player_id": 999
}
```

#### Get Players by Position

```bash
GET http://localhost:3000/player/position/1
```

**Expected:** All Forward players (Messi, Ronaldo, Mbappé)

#### Fetch User Team

```bash
GET http://localhost:3000/player/user-team
Headers:
  user-id: 1
```

---

## 🔧 Testing Scenarios

### Scenario 1: User Profile with Full Joins

1. Call `GET /user/profile` with `user-id: 1`
2. Verify response includes:
    - User basic info
    - `best_player` object with Messi's details
    - `best_club` object with Barcelona's details
    - `best_manager` object with Xavi's details
    - `best_country` object with Spain's details

### Scenario 2: Catalog CRUD Operations

1. List all clubs: `GET /catalog/clubs`
2. Get Barcelona: `GET /catalog/clubs/1`
3. Update Barcelona: `PUT /catalog/clubs/1`
4. Create new club: `POST /catalog/clubs`
5. Delete club: `DELETE /catalog/clubs/{new_id}`

### Scenario 3: User Interactions

1. Get user 1's players: `GET /user/players` (header: user-id: 1)
2. Add new players to user 1: `POST /user/players`
3. Make user 1 follow user 2: `POST /user/follow`
4. Make user 1 unfollow user 2: `POST /user/unfollow`

### Scenario 4: Player Management

1. List players by position: `GET /player/position/1` (Forwards)
2. Get specific player: `GET /player/1` (Messi)
3. Create new player: `POST /player/`
4. Update player: `PUT /player/{id}`

---

## 🐛 Troubleshooting

### Common Issues:

1. **"relation does not exist" errors:**
    - Ensure migrations ran successfully
    - Check database connection in `config/local.yaml`

2. **"record not found" errors:**
    - Verify mock data was inserted (migration 0011)
    - Check the ID values in your requests

3. **Server won't start:**
    - Make sure port 3000 is not in use
    - Check PostgreSQL is running
    - Verify config/local.yaml database settings

### Verify Mock Data:

```sql
-- Connect to database and check:
SELECT COUNT(*) FROM users;     -- Should be 5
SELECT COUNT(*) FROM players;   -- Should be 12  
SELECT COUNT(*) FROM clubs;     -- Should be 6
```

---

## 📋 Postman Collection

Import the updated Postman collection from:
`docs/postman/FanligaFc.postman_collection.json`

The collection includes:

- ✅ All catalog CRUD operations
- ✅ Complete user management
- ✅ Player management APIs
- ✅ Pre-configured environment variables
- ✅ Sample request bodies with mock data IDs

**Environment Variables:**

- `baseUrl`: http://localhost:3000
- `userId`: 1 (or any user ID 1-5)
- `playerId`: 1 (or any player ID 1-12)
- `clubId`: 1 (or any club ID 1-6)
- And more...

---

## 🎯 Expected Results

With the comprehensive mock data, you should be able to:

1. **Test User Profiles** - See complete nested data with joins
2. **CRUD All Catalog Items** - Positions, Clubs, Managers, Countries, Players, Entities
3. **User Interactions** - Follow/unfollow, add/remove players
4. **Player Management** - Full player lifecycle operations
5. **Data Relationships** - Test foreign key relationships and constraints

The mock data provides realistic football/soccer data with proper relationships between users,
players, clubs, managers, and countries for comprehensive API testing.