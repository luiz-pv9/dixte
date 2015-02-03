-- APPS ------------------------------------------------------------------------
CREATE SEQUENCE apps_id_seq;
CREATE TABLE apps (
	app_id  BIGINT NOT NULL DEFAULT NEXTVAL('apps_id_seq'),
	name VARCHAR(80) NOT NULL,
	token VARCHAR(80) NOT NULL
);
ALTER SEQUENCE apps_id_seq OWNED BY apps.app_id;
ALTER TABLE apps ADD PRIMARY KEY(app_id);
CREATE UNIQUE INDEX token_unique_index ON apps(token);

-- EVENTS ----------------------------------------------------------------------
CREATE SEQUENCE events_id_seq;
CREATE TABLE events (
	event_id      BIGINT NOT NULL DEFAULT NEXTVAL('events_id_seq'),
	type          VARCHAR(80) NOT NULL,
	external_id   VARCHAR(80) NOT NULL,
	happened_at   TIMESTAMP NOT NULL DEFAULT NOW(),
	properties    JSON,
	app_token     VARCHAR(80) NOT NULL
);

ALTER SEQUENCE events_id_seq OWNED BY events.event_id;
ALTER TABLE events ADD PRIMARY KEY(event_id);
ALTER TABLE events ADD CONSTRAINT app_token_fk
	FOREIGN KEY (app_token) REFERENCES apps(token)
	ON UPDATE CASCADE ON DELETE CASCADE;
CREATE INDEX app_token_type_index ON events(app_token, type);
CREATE INDEX app_token_happened_at_index ON events(app_token, happened_at);


-- PROFILES --------------------------------------------------------------------
CREATE SEQUENCE profiles_id_seq;
CREATE TABLE profiles (
	profile_id  BIGINT NOT NULL DEFAULT NEXTVAL('profiles_id_seq'),
	external_id VARCHAR(80),
	created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
	properties  JSON,
	app_token   VARCHAR(80) NOT NULL
);

ALTER SEQUENCE profiles_id_seq OWNED BY profiles.profile_id;
ALTER TABLE profiles ADD PRIMARY KEY(profile_id);
ALTER TABLE profiles ADD CONSTRAINT app_token_fk
	FOREIGN KEY (app_token) REFERENCES apps(token)
	ON UPDATE CASCADE ON DELETE CASCADE;
CREATE INDEX app_token_index ON profiles(app_token);
