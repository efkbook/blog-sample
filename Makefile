ENV ?= development
ESHOST ?= http://localhost:9200

docker_compose = $(shell which docker-compose)
scaneo = $(GOPATH)/bin/scaneo
sql-migrate = $(GOPATH)/bin/sql-migrate
godep = $(GOPATH)/bin/godep

$(scaneo):
	go get github.com/variadico/scaneo

$(sql-migrate):
	go get github.com/rubenv/sql-migrate/...

$(godep):
	go get github.com/tools/godep

deps: $(scaneo) $(sql-migrate) $(godep)

test:
	go test -v ./...

run: model/type_scans.go
	$(docker_compose) up -d

logs:
	$(docker_compose) logs

gen: model/type_scans.go

model/type_scans.go: model/type.go $(scaneo)
	cd model && go generate

app/build: $(godep)
	$(godep) restore
	$(godep) go build

app/run: $(godep)
	$(godep) restore
	$(godep) go run main.go

migrate/up: $(sql-migrate)
	$(sql-migrate) up -env=$(ENV)

elasticsearch/mapping:
	curl -XPUT "$(ESHOST)/article" -d @_elasticsearch/article.mapping.json

elasticsearch/template:
	curl -XPUT "$(ESHOST)/_template/nginx_logs_template" -d @_elasticsearch/nginx.mapping.template.json
	curl -XPUT "$(ESHOST)/_template/search_logs_template" -d @_elasticsearch/search.mapping.template.json

.PHONY: test \
	run \
	logs\
	app/run \
	migrate/up \
	elasticsearch/mapping \
	elasticsearch/template

help:
	@echo deps:                   Install dependent tools.
	@echo test:                   Run tests.
	@echo run:                    Start containers on local docker.
	@echo logs:                   Display containers logs.
	@echo gen:                    Update auto-generated code.
	@echo server/build:           Build server binary.
	@echo server/run:             Run server process.
	@echo migrate/up:             Execute database shema migration.
	@echo elasticsearch/mapping:  Register elasticseach mappings.
	@echo elasticsearch/template: Register elasticsearch mapping template.
