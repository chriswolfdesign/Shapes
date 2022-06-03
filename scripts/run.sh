#! /bin/bash

docker build -t shapes-run -f Dockerfiles/Dockerfile-run .
docker run --rm --name shapes-run shapes-run