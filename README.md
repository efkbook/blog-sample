# blog sample

[![Build Status](https://travis-ci.org/efkbook/blog-sample.svg?branch=master)](https://travis-ci.org/efkbook/blog-sample)

A kind of blog for presentating EFK stack features.

* Write, view and search blog posts.
* Sing up, Login. Session is based on cookie.
* Collect page views and search logs into Elasticsearch via fluentd in Real-Time.
* Of course, you can view and analyze user actions in Elasticsearch and Kibana.

This repository includes following.

* [fluent/fluentd](https://github.com/fluent/fluentd) v0.14 latest
* [uken/fluent-plugin-elasticsearch](https://github.com/uken/fluent-plugin-elasticsearch)
* Elastic v5 (Elasticsearch and Kibana)
* Simple blog application based on [suzuken/wiki](https://github.com/suzuken/wiki) written in Go.

## Prerequisite

* Docker

For local development,

* Go (>= 1.8)
* sqlite3

## Using Docker: make run

If you using docker, it's easy to work with Elastic stack and fluentd!

    # running docker containers by `docker-compose up -d`
    make run

At the first time, `docker-compose` start creating containers. After starting containers, it's time to access Elasticsearch. Cnfirmed by `curl`.

    curl http://localhost:9200

It works! And your Kibana console is also available on `http://localhost:5601`. If Elasticsearch is accecible, blog application can start up.

## For contributor

At first, you should install dependencies to development.

    make deps

To start a blog application locally, you just `go run`

    go run main.go

To add some external packages, use `godep`. If you want to use Elasticsearch and/or Fluentd on docker container, you can specify each host via flag. When blog app running on container, Elasticsearch and Fluentd are accessed by using docker links.

## Acknowledgement

UI Template is based on [BlackrockDigital/startbootstrap-blog-post](https://github.com/BlackrockDigital/startbootstrap-blog-post).

## LICENSE

MIT

## Author

Kenta Suzuki (a.k.a. suzuken)
