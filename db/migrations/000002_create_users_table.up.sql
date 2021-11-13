-- CREATE TABLE IF NOT EXISTS users(
-- 	id SERIAL PRIMARY KEY,
-- 	name VARCHAR(100) NOT NULL,
-- 	description TEXT,
-- 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE IF NOT EXISTS users(
	id VARCHAR(36) NOT NULL,
	email VARCHAR(225) NOT NULL unique,
	username VARCHAR(225),
	firstname VARCHAR(225),
	lastname VARCHAR(225),
	password VARCHAR(225) NOT NULL,
	tokenhash VARCHAR(15) NOT NULL,
	createdat TIMESTAMP NOT NULL,
	updatedat TIMESTAMP NOT NULL,
	primary KEY (id)
)