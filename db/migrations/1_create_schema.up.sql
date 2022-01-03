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
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	username VARCHAR(100) NOT NULL UNIQUE,
	roles VARCHAR(255),
	-- roles VARCHAR(255) NOT NULL,
	password VARCHAR(1024) NOT NULL,
	created TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create Items table
CREATE TABLE IF NOT EXISTS items(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	ownerId INT REFERENCES users(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Games table
CREATE TABLE IF NOT EXISTS games(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	ownerId INT REFERENCES users(id),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create GamePlayer tables (join table to link players to games??)
CREATE TABLE IF NOT EXISTS players(
	id SERIAL PRIMARY KEY,
	gameId INT REFERENCES games(id),
	playerId INT REFERENCES users(id)
);