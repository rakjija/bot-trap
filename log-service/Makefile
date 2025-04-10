# Makefile

SHELL := /bin/bash
DOCKER_COMPOSE := docker compose

up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down

restart: down up

logs:
	$(DOCKER_COMPOSE) logs -f

clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

build:
	docker build -t log-service ./