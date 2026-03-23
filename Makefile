.PHONY: migrate-up migrate-down migrate-up-service migrate-down-service migrate-status-service migrate-create build build-service run-service wire wire-service infra-up infra-down proto

SERVICES=user-service restaurant-service menu-service order-service review-service

# ─── Migration ───────────────────────────────────────────────────────────────

migrate-up:
	@for service in $(SERVICES); do \
		echo "→ $$service migrate up"; \
		$(MAKE) -C services/$$service migrate-up; \
	done

migrate-down:
	@for service in $(SERVICES); do \
		echo "→ $$service migrate down"; \
		$(MAKE) -C services/$$service migrate-down; \
	done

migrate-up-service:
	$(MAKE) -C services/$(svc) migrate-up

migrate-down-service:
	$(MAKE) -C services/$(svc) migrate-down

migrate-status-service:
	$(MAKE) -C services/$(svc) migrate-status

migrate-create:
	$(MAKE) -C services/$(svc) migrate-create name=$(name)

# ─── Build ───────────────────────────────────────────────────────────────────

build:
	@for service in $(SERVICES); do \
		echo "→ $$service build"; \
		$(MAKE) -C services/$$service build; \
	done

build-service:
	$(MAKE) -C services/$(svc) build

# ─── Run ────────────────────────────────────────────────────────────────────

run-service:
	$(MAKE) -C services/$(svc) run

# ─── Wire ────────────────────────────────────────────────────────────────────

wire:
	@for service in $(SERVICES); do \
		echo "→ $$service wire"; \
		$(MAKE) -C services/$$service wire; \
	done

wire-service:
	$(MAKE) -C services/$(svc) wire

# ─── Infra ───────────────────────────────────────────────────────────────────

infra-up:
	docker compose -f docker-compose.infra.yml up -d

infra-down:
	docker compose -f docker-compose.infra.yml down

# ─── Proto ───────────────────────────────────────────────────────────────────

proto:
	buf generate