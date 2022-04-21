-- Insert [User] Mr Sir
INSERT INTO "users" (id, name, username, password)
	VALUES (
		'TPEoMpD1A4hAfi0VTGjmb',
		'Mr Sir',
		'mrsir',
		'$2a$04$P4ouhozPJZX8NCCm7QyrIe1ZR46HNKL5tZgr0Yn4RCPyY85hnAM0m'
	);

-- Insert [User] Fred
INSERT INTO "users" (id, name, username, password)
	VALUES (
		'nF9Dvdr28Ds-XbeQib4Nz',
		'fred',
		'fredo',
		'$2a$04$7ivgCT/hC7I24V0idCEppe8271qodQRtWVeMN/R9KtOS3iXo9qduW'
	);


INSERT INTO "items" (name, description, owner_id)
	VALUES ('Item #1', 'Item #1 Created From Seed Script', 'TPEoMpD1A4hAfi0VTGjmb');

INSERT INTO "items" (name, description, owner_id)
	VALUES ('Item #2', 'Item #2 Created From Seed Script', 'TPEoMpD1A4hAfi0VTGjmb');


INSERT INTO "games" (id, name, description, owner_id, share_id)
	VALUES (
		'TPEoMpD1A4hAfi0VTGjma',
		'SEED: First Game',
		'Game Description',
		'TPEoMpD1A4hAfi0VTGjmb',
		'TPEoMpD1A4hAfi0VTGyyy'
	);


INSERT INTO "players" (id, game_id, player_id)
	VALUES (
		'TPEoMpD1A4hAfi0VTGxxx',
		'TPEoMpD1A4hAfi0VTGjma',
		'TPEoMpD1A4hAfi0VTGjmb'
	)