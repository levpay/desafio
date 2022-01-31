up-db:
	docker-compose up db

create-db-schema:
	chmod 755 ./create_database.sh
	./create_database.sh
