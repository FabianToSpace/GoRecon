#!/bin/bash
CGO_ENABLED=0 go build -gcflags "all=-N -l" -o gorecon .
(docker compose -f ./docker/debug.docker-compose.yaml up gorecon -d --build)
sleep 5