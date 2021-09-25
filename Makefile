.PHONY: start
start:
	@./bin/start
	./bin/migrate

.PHONY: migrate
migrate:
	@./bin/migrate
