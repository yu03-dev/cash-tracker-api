build:
	@docker compose build

up:
	@docker compose up

up_d:
	@docker compose up -d

down:
	@docker compose down

migrate_up:
	@docker exec -it cash-tracker-api ./bin/migrate up

migrate_down:
	@docker exec -it cash-tracker-api ./bin/migrate down

exec_db:
	@docker exec -it postgres-db psql -d postgres -U postgres