.PHONY: all build

BIN_DIR := ./bin
version := $(shell git rev-parse --short=12 HEAD)
timestamp := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
dbdir=/tmp/nftmaven/db
dbinitdir=$(ROOT_DIR)/deployments/db


all: build


build:
	rm -f $(BIN_DIR)/osc
	go build -o $(BIN_DIR)/osc -v -ldflags "-X main.rev=$(version) -X main.bts=$(timestamp)" cmd/osc/main.go

dockerinit:
	-docker container prune -f >/dev/null 2>&1
	-docker network create nftmavennet >/dev/null 2>&1


dbinit: dbhalt
	-docker container prune -f >/dev/null 2>&1
	-sudo rm -rf $(dbdir)
	-docker run --detach -v $(dbdir):/var/lib/mysql:z  -v $(dbinitdir):/docker-entrypoint-initdb.d:z --network nftmavennet --name nftmavendb --env MARIADB_USER=$(NFTMAVEN_DB_USER) --env MARIADB_PASSWORD=$(NFTMAVEN_DB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(NFTMAVEN_DB_ROOT_PASSWORD) --env MARIADB_DATABASE=$(NFTMAVEN_DB_DATABASE) mariadb:latest --port 13306


dbstart:
	-docker container prune -f >/dev/null 2>&1
	-docker run --detach -v $(dbdir):/var/lib/mysql  --network nftmavennet --name nftmavendb --env MARIADB_USER=$(NFTMAVEN_DB_USER) --env MARIADB_PASSWORD=$(NFTMAVEN_DB_PASSWORD) --env MARIADB_ROOT_PASSWORD=$(NFTMAVEN_DB_ROOT_PASSWORD) mariadb:latest --port 13306


dbhalt:
	-docker stop nftmavendb
	-docker container prune -f >/dev/null 2>&1

dbprompt:
	-docker container prune -f >/dev/null 2>&1
	-docker run --network nftmavennet -it --rm mariadb mysql -h nftmavendb -u $(NFTMAVEN_DB_USER) -D $(NFTMAVEN_DB_DATABASE) -p$(NFTMAVEN_DB_PASSWORD) --port 13306
