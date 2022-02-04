#!/bin/bash
DB_USER="placeholder-user"
DB_HOST="localhost"
DB_NAME="challenge"

psql -U ${DB_USER} -h ${DB_HOST} -d ${DB_NAME} -c \
    'CREATE TABLE supers(
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
    );'