DOCKER_COMPOSE=$(shell which docker-compose)

deps:

run:
	$(DOCKER_COMPOSE) up

logs:
	$(DOCKER_COMPOSE) logs
