CREATE DATABASE IF NOT EXISTS superhero_api;

CREATE TABLE supers(
    id SERIAL,
    uuid VARCHAR NOT NULL,
    hero_name VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    alignment VARCHAR NOT NULL,
    intelligence VARCHAR(10) NOT NULL,
    power VARCHAR(10) NOT NULL,
    occupation VARCHAR NOT NULL,
    image VARCHAR NOT NULL,
    group_connections VARCHAR,
    relatives VARCHAR
);

