DOCKER_COMPOSE=$(shell which docker-compose)
ENV:=development
ESHOST:=http://localhost:9200

deps:
	which sql-migrate || go get github.com/rubenv/sql-migrate/...
	which scaneo || go get github.com/variadico/scaneo
	which godep || go get github.com/tools/godep

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

elasticsearch/mapping:
	curl -XPUT "$(ESHOST)/article" -d @_elasticsearch/article.mapping.json

elasticsearch/template:
	curl -XPUT "$(ESHOST)/_template/nginx_logs_template" -d @_elasticsearch/nginx.mapping.template.json
	curl -XPUT "$(ESHOST)/_template/search_logs_template" -d @_elasticsearch/search.mapping.template.json
