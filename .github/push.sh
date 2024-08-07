#!/bin/bash

rm .env

REPOSITORY="${DOCKER_USERNAME}/medzoner-go"
REPOSITORY_NEW="${REPOSITORY}:${GIT_BRANCH}-${GIT_COMMIT}"

echo "Registry login..."
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin

echo "Building image ${REPOSITORY} ..."
docker build -t ${REPOSITORY} -f ./docker/Dockerfile .

echo "Tag image ${REPOSITORY_NEW} ..."
docker tag "${REPOSITORY}" "${REPOSITORY_NEW}"

echo "Pushing image to registry ${REPOSITORY_NEW} ..."
docker push "${REPOSITORY_NEW}"
