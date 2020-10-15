# network-service

A web app which allows you to see all hosts in your network.

Uses:

* HTML templates
* Redis

## how to start

* Run `docker-compose up --build` to start redis.
* Run `REDIS_HOST=localhost REDIS_PORT=6379 TARGET=${ROUTER_HOST} go run cmd/main.go`