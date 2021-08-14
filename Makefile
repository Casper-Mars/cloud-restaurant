.PHONY: set-up-dev
set-up-dev:
	docker-compose -f ./hack/env/docker-compose.yml up -d

.PHONY: stop-dev
stop-dev:
	docker-compose -f ./hack/env/docker-compose.yml down
