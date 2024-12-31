build: 
	docker compose build
.PHONY: build
up: 
	docker compose up --attach app
.PHONY: up

app: docker compose up --attach app --build
.PHONY: app