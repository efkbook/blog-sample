DOCKER_COMPOSE=$(shell which docker-compose)
ENV:=development

deps:
	which sql-migrate || go get github.com/rubenv/sql-migrate/...
	which scaneo || go get github.com/variadico/scaneo

test:
	go test -v ./...

run:
	$(DOCKER_COMPOSE) up -d

logs:
	$(DOCKER_COMPOSE) logs

migrate/up:
	sql-migrate up -env=$(ENV)

gen:
	cd model && go generate
