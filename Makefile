COMPOSE_FILE = docker-compose.yml
DOCKER_COMPOSE = docker-compose -f $(COMPOSE_FILE)

.PHONY: help run install-prometheus down-prometheus clean

help:
	@echo "Comandos disponíveis:"
	@echo "  make run                - Executar projeto"
	@echo "  make install-prometheus - Subir os serviços do Prometheus"
	@echo "  make down-prometheus    - Derrubar os serviços do Prometheus"
	@echo "  make clean              - Remover recursos Docker criados"

run:
	air

install-prometheus:
	$(DOCKER_COMPOSE) up -d

down-prometheus:
	$(DOCKER_COMPOSE) down

clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans
	docker image prune -f

build: 
	go mod tidy
	set GOARCH=amd64 && set GOOS=linux && go build -v -ldflags="-s -w" -o bin/main ./main.go
