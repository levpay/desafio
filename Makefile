up-db:
	docker-compose up db

up-api:
	docker-compose up api

create-db-schema:
	chmod 755 ./utils/create_database.sh
	./utils/create_database.sh

build:
	docker build -t challenge:latest .

test:
	go test -v ./...