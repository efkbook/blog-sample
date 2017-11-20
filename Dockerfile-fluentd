FROM debian:jessie

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get -qq update && apt-get install --no-install-recommends -y curl ca-certificates sudo build-essential libcurl4-gnutls-dev
RUN curl -L https://toolbelt.treasuredata.com/sh/install-debian-jessie-td-agent2.sh | sh
RUN /usr/sbin/td-agent-gem install fluent-plugin-elasticsearch fluent-plugin-record-reformer

EXPOSE 24224

CMD exec td-agent -c /fluentd/etc/$FLUENTD_CONF -p /fluentd/plugins $FLUENTD_OPT
