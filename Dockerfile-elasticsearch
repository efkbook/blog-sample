# based on images from docker.elastic.co
# see https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html
FROM docker.elastic.co/elasticsearch/elasticsearch:5.4.3

# Install plugins
RUN \
  cd /usr/share/elasticsearch && \
  bin/elasticsearch-plugin install analysis-kuromoji && \
  bin/elasticsearch-plugin install analysis-icu
