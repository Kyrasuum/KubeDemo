.PHONY: run
#: Starts the project
run: .deps
	@docker compose --profile all up --pull never -d

.PHONY: stop
#: Stops the project
stop: logs
	@docker compose --profile all down

.PHONY: restart
#: Restarts the project
restart: stop run

.PHONY: rebuild
#: Creates a fresh build of the project
rebuild: .deps stop
	@docker compose build rtr
	@docker compose build web
	@docker compose build api

.PHONY: build
#: Performs a clean run of the project
build: rebuild run

.PHONY: web
#: Builds and runs web
web:
	$(MAKE) --no-print-directory -C web web

.PHONY: web
#: Builds and runs api
web:
	$(MAKE) --no-print-directory -C api api

.PHONY: logs
#: Extracts and saves logs locally from running cerebro containers
logs:
	@images=`docker ps -a | awk '{print $$NF}' | grep -G "^kubedemo-"`; \
	echo $$images | tr " " "\n" | awk '{ system("docker logs " $$1 "-1 2> logs/" $$1 ".log > logs/" $$1 ".log") }' >/dev/null

.PHONY: deps
.deps:
	$(MAKE) --no-print-directory deps

#: Install dependencies for docker and targets in this makefile
deps:
	@if ! command -v docker; then wget https://get.docker.com -O - | sh; fi
	@if ! cat /etc/group | grep docker; then sudo groupadd docker; fi
	@sudo usermod -aG docker $(USER)
	@sudo apt install docker-compose-plugin
	@sudo touch /etc/docker/daemon.json
	@touch .deps

.PHONY: dev-deps
.dev-deps:
	@$(MAKE) --no-print-directory dev-deps

#: Installs all depedencies for development
dev-deps: .deps
	@wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
	@export NVM_DIR="$$([ -z "$${XDG_CONFIG_HOME-}" ] && printf %s "$${HOME}/.nvm" || printf %s "$${XDG_CONFIG_HOME}/nvm")" \
		&& [ -s "$$NVM_DIR/nvm.sh" ] && \. $$NVM_DIR/nvm.sh \
		&& nvm install v18.6.0 && nvm use v18.6.0 \
		&& nvm install-latest-npm \
		&& cd web/internal/pages && npm run clean && npm clean-install
	@sudo apt-add-repository --update ppa:longsleep/golang-backports
	@sudo apt install golang-1.22 postgresql-client-14
	@sudo apt install pip
	@touch .dev-deps

.PHONY: clean
#: Cleans slate for docker images
clean: stop
	@docker container prune -f
	@docker system prune -f
	@docker image rm $$(docker images -a | grep "kubedemo" | grep -v "alpine" | awk '{print $$3}') --force || true;

cl:
	@make clean && make clean && make clean-vol

.PHONY: help
#: Lists available commands
help:
	@echo "Available Commands for project:"
	@grep -B1 -E "^[a-zA-Z0-9_-]+\:([^\=]|$$)" Makefile \
	 | grep -v -- -- \
	 | sed 'N;s/\n/###/' \
	 | sed -n 's/^#: \(.*\)###\(.*\):.*/\2###\1/p' \
	 | column -t  -s '###'