INSERT INTO "users" (name, username, password)
	VALUES (
		'Mr Sir',
		'mrsir',
		'$2a$04$P4ouhozPJZX8NCCm7QyrIe1ZR46HNKL5tZgr0Yn4RCPyY85hnAM0m'
	);

INSERT INTO "items" (name, description, owner_id)
	VALUES ('Item #1', 'Item #1 Created From Seed Script', 1);

INSERT INTO "items" (name, description, owner_id)
	VALUES ('Item #2', 'Item #2 Created From Seed Script', 1);

INSERT INTO "games" (name, description, owner_id, share_id)
	VALUES ('First Game', 'Game Description', 1, 'TPEoMpD1A4hAfi0VTGjmb');

INSERT INTO "players" (game_id, player_id)
	VALUES (1, 1)