#! /bin/bash

docker build -t shapes-test -f Dockerfiles/Dockerfile-test .
docker run --rm --name shapes-test shapes-test