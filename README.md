# GOBB

## What?
GoBB is a BB forum written in Golang

## Why?
This is just an exercise in Go, using dependency injection via Wire and Docker to run the service dependencies such as Postgres.
I started my development career 12 years ago writing forums in PHP, Ruby and even Node so it only seems fitting to write one in Go.

## How?
Start the app with `make start`. This runs the app in a Docker container and starts up all dependencies.

## Migrations

* Run migrations with `make migrate`
* Create migrations using `migrate`
    * `migrate create -ext sql -dir db/migrations create_topics_table`
