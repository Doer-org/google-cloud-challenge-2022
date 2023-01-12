
.DEFAULT_GOAL := help

# env
ENV_DEV_FILE := .env.local

CLIENT_SRC_DIR   := ./src

# docker
DOCKER_COMPOSE_DIR                 := ./docker
DOCKER_COMPOSE_DEV                 := $(DOCKER_COMPOSE_DIR)/docker-compose.yml

DOCKER_EXEC                        := docker exec -it
DOCKER_COMPOSE_PJ_NAME             := gc-challenge-2022
CLIENT_CONTAINER_NAME              := $(DOCKER_COMPOSE_PJ_NAME)-client


NODEMODULES_DIR := $(CLIENT_SRC_DIR)/node_modules

# rm
RM := rm -rf

.PHONY: up
up: ## docker環境を立ち上げる
	docker-compose \
	-f $(DOCKER_COMPOSE_DEV) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	up -d

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose \
	-f $(DOCKER_COMPOSE_DEV) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	down \
	--rmi all --volumes --remove-orphans

.PHONY: fclean
fclean:down del-volumes ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:del-nodemodules

.PHONY: del-nodemodules
del-nodemodules:
	$(RM) $(NODEMODULES_DIR)

.PHONY: log
log: ## docker compose環境のログを確認する
	docker-compose \
	-f $(DOCKER_COMPOSE_DEV) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	logs -f

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
