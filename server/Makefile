
.DEFAULT_GOAL := help

# env
ENV_DEV_FILE := .env.dev

# go
GO           := go run
GO_SRC_DIR   := ./src

# docker
DOCKER_COMPOSE_DIR                 := ./docker
DOCKER_COMPOSE_DEV_DIR             := $(DOCKER_COMPOSE_DIR)/dev
DOCKER_COMPOSE_DEV                 := $(DOCKER_COMPOSE_DEV_DIR)/docker-compose.yml
DOCKER_COMPOSE_DEV_DB              := $(DOCKER_COMPOSE_DEV_DIR)/docker-compose.db.yml
DOCKER_COMPOSE_DEV_SERVER          := $(DOCKER_COMPOSE_DEV_DIR)/docker-compose.server.yml
DOCKER_EXEC                        := docker exec -it
DOCKER_COMPOSE_PJ_NAME             := gc-challenge-2022-api
SERVER_CONTAINER_NAME              := $(DOCKER_COMPOSE_PJ_NAME)-server
DB_CONTAINER_NAME                  := $(DOCKER_COMPOSE_PJ_NAME)-db
DOCKER_COMPOSE_DEFAULT_OPTIONS     := -f $(DOCKER_COMPOSE_DEV) \
									-f $(DOCKER_COMPOSE_DEV_SERVER) \
									-f $(DOCKER_COMPOSE_DEV_DB) \
									-p $(DOCKER_COMPOSE_PJ_NAME)
# volume
DATA_DIR := ./db/data


#ent
ENT_DIR := $(GO_SRC_DIR)/infrastructure/ent

# rm
RM := rm -rf

.PHONY: up
up: ent-generate ## docker環境を立ち上げる
	docker-compose $(DOCKER_COMPOSE_DEFAULT_OPTIONS) \
	up -d

.PHONY: down
down: ## dockerイメージを削除し、docker環境を閉じる
	docker-compose $(DOCKER_COMPOSE_DEFAULT_OPTIONS) \
	down \
	--rmi all --volumes --remove-orphans

.PHONY: fclean
fclean:down del-volumes ## マウントしたデータを削除、またdockerイメージも削除する

.PHONY: re
re:fclean up ## 完全に初期化した状態でdocker環境を立ち上げる

.PHONY: del-volumes
del-volumes:del-data

.PHONY: del-data
del-data:
	$(RM) $(DATA_DIR)

.PHONY: attach-server
attach-server: ## dockerのserverコンテナにアクセスする
	$(DOCKER_EXEC) $(SERVER_CONTAINER_NAME) bash

.PHONY: attach-db
attach-db: ## dockerのdbコンテナにアクセスする
	$(DOCKER_EXEC) $(DB_CONTAINER_NAME) bash

.PHONY: server-lint
go-lint: ## src配下のコードを整形する
	cd $(GO_SRC_DIR) && gofmt -l -w .

.PHONY: test
test: ## テストを実行する
	cd src/ && go test -v ./...

.PHONY: log
log: ## docker compose環境のログを確認する
	docker-compose $(DOCKER_COMPOSE_DEFAULT_OPTIONS) \
	logs -f

#ent
ent-generate: ## entのファイルを自動生成する
	cd $(ENT_DIR) && go generate .

ent-describe: ## entのスキーマの詳細を表示する
	cd $(ENT_DIR) && go run -mod=mod entgo.io/ent/cmd/ent describe ./schema

.PHONY: help
help: ## コマンドの一覧を標示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'