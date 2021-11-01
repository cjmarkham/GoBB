.PHONY: start
start:
	@./bin/start

.PHONY: stop
stop:
	@./bin/stop

.PHONY: migrate
migrate:
	./bin/migrate

.PHONY: seed
seed:
	@./bin/seed
