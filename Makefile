include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@CMD="go run ./cmd/api"; \
	[[ -n "$(SERVICE_NAME)" ]] && CMD="$$CMD -name='$(SERVICE_NAME)'"; \
	[[ -n "$(SERVICE_DESC)" ]] && CMD="$$CMD -desc='$(SERVICE_DESC)'"; \
	[[ -n "$(ENV)" ]] && CMD="$$CMD -env='$(ENV)'"; \
	[[ -n "$(REGISTRY_ADDR)" ]] && CMD="$$CMD -registry='$(REGISTRY_ADDR)'"; \
	eval $$CMD

## run/gateway: run the cmd/gateway application
.PHONY: run/gateway
run/gateway:
	@CMD="go run ./cmd/gateway"; \
	[[ -n "$(GATEWAY_IP)" ]] && CMD="$$CMD -ip='$(GATEWAY_IP)'"; \
	[[ -n "$(GATEWAY_PORT)" ]] && CMD="$$CMD -port='$(GATEWAY_PORT)'"; \
	[[ -n "$(REDIS_HOST)" ]] && CMD="$$CMD -redisHost='$(REDIS_HOST)'"; \
	[[ -n "$(REDIS_PORT)" ]] && CMD="$$CMD -redisPort='$(REDIS_PORT)'"; \
	[[ -n "$(REDIS_PASSWORD)" ]] && CMD="$$CMD -redisPassword='$(REDIS_PASSWORD)'"; \
	eval $$CMD

## run/registry: run the cmd/registry application
.PHONY: run/registry
run/registry:
	@CMD="go run ./cmd/registry"; \
	[[ -n "$(REGISTRY_IP)" ]] && CMD="$$CMD -ip='$(REGISTRY_IP)'"; \
	[[ -n "$(REGISTRY_PORT)" ]] && CMD="$$CMD -port='$(REGISTRY_PORT)'"; \
	[[ -n "$(ENV)" ]] && CMD="$$CMD -env='$(ENV)'"; \
	[[ -n "$(CLEANUP_INTERVAL)" ]] && CMD="$$CMD -cleanup='$(CLEANUP_INTERVAL)'"; \
	[[ -n "$(REDIS_HOST)" ]] && CMD="$$CMD -redisHost='$(REDIS_HOST)'"; \
	[[ -n "$(REDIS_PORT)" ]] && CMD="$$CMD -redisPort='$(REDIS_PORT)'"; \
	[[ -n "$(REDIS_PASSWORD)" ]] && CMD="$$CMD -redisPassword='$(REDIS_PASSWORD)'"; \
	eval $$CMD

## up/redis: run the docker compose redis file 
.PHONY: up/redis
up/redis:
	docker compose -f docker-compose-redis.yml up -d

## down/redis: stop the docker containers for redis file 
.PHONY: down/redis
down/redis:
	docker compose -f docker-compose-redis.yml down

