#!/bin/bash
DB_USER="placeholder-user"
DB_HOST="localhost"
DB_NAME="challenge"

psql -U ${DB_USER} -h ${DB_HOST} -d ${DB_NAME} -c \
    'CREATE TABLE supers(
        id SERIAL,
        uuid VARCHAR(255) NOT NULL,
        hero_name VARCHAR(255) NOT NULL,
        intelligence VARCHAR(10) NOT NULL,
        power VARCHAR(10) NOT NULL,
        occupation VARCHAR(255) NOT NULL,
        image VARCHAR(255) NOT NULL,
        group_connections VARCHAR(255),
        relatives VARCHAR(255)
    );'