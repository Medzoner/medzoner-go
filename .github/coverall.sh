#!/bin/bash

if [ "${GIT_BRANCH}" = "master" ]; then
  goveralls -coverprofile=coverage.out -service=travis-ci -repotoken "${COVERALLS_TOKEN}";
fi
