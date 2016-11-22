DOCKER_COMPOSE=$(shell which docker-compose)
ENV:=development

deps:

run:
	$(DOCKER_COMPOSE) up

logs:
	$(DOCKER_COMPOSE) logs

migrate/up:
	sql-migrate up -env=$(ENV)
