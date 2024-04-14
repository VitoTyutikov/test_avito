docker = docker-compose

all:run

run:
	$(docker) up server -d

stop:
	$(docker) down

restart: stop run

test:
	$(docker) up tester

load_test:run
	go run ./tests/load/load.go