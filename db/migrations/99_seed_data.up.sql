INSERT INTO "users" (name, username, password)
	VALUES ('Mr Sir', 'mrsir', '$2a$04$P4ouhozPJZX8NCCm7QyrIe1ZR46HNKL5tZgr0Yn4RCPyY85hnAM0m');

INSERT INTO "items" (name, description, ownerId)
	VALUES ('Item #1', 'Item #1 Created From Seed Script', 1);

INSERT INTO "items" (name, description, ownerId)
	VALUES ('Item #2', 'Item #2 Created From Seed Script', 1);

INSERT INTO "games" (name, description, ownerId, share_id)
	VALUES ('First Game', 'Game Description', 1, 'unique id #1');

INSERT INTO "players" (gameId, playerId)
	VALUES (1, 1)