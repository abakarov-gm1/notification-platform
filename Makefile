
up-auth:
	docker compose run --rm -p 8081:8081 auth-service /bin/bash

up-fake:
	docker compose run --rm -p 8084:8081 fake-kafka-service /bin/bash

up:
	docker compose up -d

down:
	docker compose down

