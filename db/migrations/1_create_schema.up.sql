-- CREATE TABLE IF NOT EXISTS users(
-- 	id SERIAL PRIMARY KEY,
-- 	-- email VARCHAR(225) NOT NULL unique,
-- 	username VARCHAR(225) NOT NULL unique,
-- 	first_name VARCHAR(225) NOT NULL,
-- 	last_name VARCHAR(225) NOT NULL,
-- 	password VARCHAR(225) NOT NULL,
-- 	token_hash VARCHAR(500) NOT NULL,
-- 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

-- Create Users table
CREATE TABLE IF NOT EXISTS users(
	id VARCHAR(200) NOT NULL UNIQUE PRIMARY KEY, -- Using VARCHAR(200) since our Ids are now using gonanoid rather than an autoincrementing number
	name VARCHAR(255) NOT NULL,
	username VARCHAR(100) NOT NULL UNIQUE,
	-- roles VARCHAR(255),
	password VARCHAR(1024) NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create Items table
CREATE TABLE IF NOT EXISTS items(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	owner_id VARCHAR(200) NOT NULL REFERENCES users(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Games table
CREATE TABLE IF NOT EXISTS games(
	id VARCHAR(200) NOT NULL UNIQUE PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	share_id VARCHAR(200) NOT NULL UNIQUE,
	owner_id VARCHAR(200) NOT NULL REFERENCES users(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create GamePlayer tables (join table to link players to games??)
CREATE TABLE IF NOT EXISTS players(
	id SERIAL PRIMARY KEY,
	game_id VARCHAR(200) NOT NULL REFERENCES games(id),
	player_id VARCHAR(200) NOT NULL REFERENCES users(id)
);