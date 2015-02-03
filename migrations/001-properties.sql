-- PROPERTIES ------------------------------------------------------------------
CREATE SEQUENCE properties_id_seq;
CREATE TABLE properties (
	property_id BIGINT NOT NULL DEFAULT NEXTVAL('properties_id_seq'),
	key VARCHAR(80) NOT NULL,
	name VARCHAR(80) NOT NULL,
	type VARCHAR(40) NOT NULL DEFAULT 'string'
);
ALTER SEQUENCE properties_id_seq OWNED BY properties.property_id;
ALTER TABLE properties ADD PRIMARY KEY(property_id);
CREATE INDEX key_properties ON properties(key);

-- VALUES ----------------------------------------------------------------------
CREATE SEQUENCE property_values_id_seq;
CREATE TABLE property_values (
	property_values_id BIGINT NOT NULL DEFAULT NEXTVAL('property_values_id_seq'),
	property_id BIGINT NOT NULL,
	value VARCHAR(120) NOT NULL,
	count BIGINT NOT NULL DEFAULT 1
);
ALTER SEQUENCE property_values_id_seq OWNED BY property_values.property_values_id;
ALTER TABLE property_values ADD PRIMARY KEY(property_values_id);
ALTER TABLE property_values ADD CONSTRAINT property_id_fk
	FOREIGN KEY (property_id) REFERENCES properties(property_id)
	ON UPDATE CASCADE ON DELETE CASCADE;

