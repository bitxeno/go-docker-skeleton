# go-docker-skeleton [![license](https://img.shields.io/badge/license-Apache%202-blue?style=flat)](https://github.com/bitxeno/go-docker-skeleton/blob/master/LICENSE)[![version](https://img.shields.io/badge/version-0.1.0-blue.svg)](https://github.com/bitxeno/go-docker-skeleton/releases)
Skeleton for run go service in docker


## Prerequisite

* Go 1.8+
* Node.js 14+
* Vue3 + daisyUI

## How To Start

1. copy all file to project directory
2. replace all `github.com/bitxeno/go-docker-skeleton` string to your repo
3. update `AppName` and `AppDesc` variables in `main.go`
4. execute shell:
```shell
cd view
npm install
cd ..
go mod vendor

# open shell #1:
# for frontend hot reload
cd view
npm run dev

# open shell #2:
go run -tags dev . server
# or hot reload
go get -u https://github.com/cosmtrek/air
air -c .air.toml
```

## How to integrated with **vue-admin-template**

```
rm -rf view
git clone --depth=1 https://github.com/PanJiaChen/vue-admin-template.git view
cd view
rm -rf ./.git
npm install
```
chnage `package.json` dev script to:
```
"dev": "vue-cli-service build --watch --mode production",
```

## How to push DockerHub

1. register dockerhub and create a repo
2. on Dockerhub, goto `Account Settings -> Security` create aceess token
3. on Github, goto repo `Settings -> Secrets` and add three github action variables
```
DOCKER_USERNAME
DOCKER_TOKEN
DOCKER_REPOSITORY
```


## Run on Docker

```
docker run -d --name=app-name --restart=always -p 8080:80  -v /path/to/config/dir:/data xxxx/app-name
```

## Great library

* [validator](https://github.com/go-playground/validator): Package validator implements value validations for structs and individual fields based on tags.
* [lo](https://github.com/samber/lo): samber/lo is a Lodash-style Go library based on Go 1.18+ Generics.
* [cron](https://github.com/robfig/cron): a cron library for go
* [gocron](https://github.com/go-co-op/gocron): Easy and fluent Go cron scheduling.
* [gjson](https://github.com/tidwall/gjson): Get JSON values quickly - JSON parser for Go
* [ristretto](https://github.com/dgraph-io/ristretto): A high performance memory-bound Go cache
* [goutil](https://github.com/gookit/goutil): Helper Utils(600+)
* [emitter](https://github.com/olebedev/emitter): Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins
* [event](https://github.com/gookit/event):  Lightweight event manager and dispatcher implements by Go
* [gods](https://github.com/emirpasic/gods): GoDS (Go Data Structures) - Sets, Lists, Stacks, Maps, Trees, Queues, and much more