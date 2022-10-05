db-up:
	docker-compose up -d db
db-down:
	docker-compose down
migrate-up: db-up
	docker-compose up migrate-up

.PHONY: db-up db-down migrate-up