#!/bin/bash
docker build -t eleven-bff:$1 -f Dockerfile . && \
docker tag eleven-bff:$1 registry.digitalocean.com/eleven-st/eleven-bff:$1 && \
docker push registry.digitalocean.com/eleven-st/eleven-bff:$1