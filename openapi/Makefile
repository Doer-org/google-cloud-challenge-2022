
.DEFAULT_GOAL := help

# docker
DOCKER_COMPOSE_DIR         := ./docker
DOCKER_COMPOSE             := $(DOCKER_COMPOSE_DIR)/docker-compose.yml
DOCKER_COMPOSE_PJ_NAME     := gc-challenge-2022-swagger

# rm
RM := rm -rf


.PHONY: up
up: ## docker環境を立ち上げる
	docker-compose \
	-f $(DOCKER_COMPOSE) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	up -d

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose \
	-f $(DOCKER_COMPOSE) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	down \
	--volumes --remove-orphans

.PHONY: fclean
fclean:down ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: log
log: ## docker compose環境のログを確認する
	docker-compose \
	-f $(DOCKER_COMPOSE_DEV) \
	-f $(DOCKER_COMPOSE_DEV_SERVER) \
	-f $(DOCKER_COMPOSE_DEV_DB) \
	-p $(DOCKER_COMPOSE_PJ_NAME) \
	logs -f

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'