#!/bin/bash

rm .env
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
docker build -t ${DOCKER_USERNAME}/medzoner-go -f ./docker/Dockerfile .
docker tag "${DOCKER_USERNAME}/medzoner-go" "${DOCKER_USERNAME}/medzoner-go:${GIT_BRANCH}-${GIT_COMMIT}"
docker push "${DOCKER_USERNAME}/medzoner-go:${GIT_BRANCH}-${GIT_COMMIT}"
