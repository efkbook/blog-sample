# blog sample

[![Build Status](https://travis-ci.org/suzuken/blog-sample.svg?branch=master)](https://travis-ci.org/suzuken/blog-sample)

A kind of blog for presentating ELK stack features.

* Write, view and search blog posts.
* Sing up, Login. Session is based on cookie.
* Collect page views and user search logs into Elasticsearch via fluentd in Real-Time.
* Of course, you can view and analyze user actions in Elasticsearch and Kibana.

This repository includes following.

* [fluent/fluentd](https://github.com/fluent/fluentd) v0.14 latest
* [uken/fluent-plugin-elasticsearch](https://github.com/uken/fluent-plugin-elasticsearch)
* Elastic v5 (Elasticsearch and Kibana)
* Simple blog application based on [suzuken/wiki](https://github.com/suzuken/wiki) writtern in Go.

## Prerequisite

* Docker
* Go (>= 1.7)
* sqlite3

## Bootstrap

    make deps

## How to run

    make run

## Acknowledgement

UI Template is based on [BlackrockDigital/startbootstrap-blog-post](https://github.com/BlackrockDigital/startbootstrap-blog-post).

## LICENSE

MIT

## Author

Kenta Suzuki (a.k.a. suzuken)
