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
* Go (>= 1.7)
* sqlite3

## Bootstrap

    make deps

## Using Docker: make run

If you using docker, it's easy to work with Elastic stack and fluentd!

    # running docker containers by docker-compose
    make run

At the first time, `docker-compose` start creating containers. After starting containers, it's time to access Elasticsearch. Cnfirmed by `curl`.

    curl http://localhost:9200

It works! And your Kibana console is also available on `http://localhost:5601`. If Elasticsearch is accecible, blog application can start up. To start a blog application,

    go run main.go

or, use `go build -o blog && ./blog`. Of course, `go get github.com/efkbook/blog-sample && blog-sample` is working as intented.

## Acknowledgement

UI Template is based on [BlackrockDigital/startbootstrap-blog-post](https://github.com/BlackrockDigital/startbootstrap-blog-post).

## LICENSE

MIT

## Author

Kenta Suzuki (a.k.a. suzuken)
