.PHONY: start
start:
	./bin/start && ./bin/migrate

.PHONY: stop
stop:
	@./bin/stop

.PHONY: migrate
migrate:
	./bin/migrate

.PHONY: seed
seed:
	@./bin/seed
