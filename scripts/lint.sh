#! /bin/bash

docker build -t shapes-lint -f Dockerfiles/Dockerfile-lint .
docker run --rm --name shapes-lint shapes-lint