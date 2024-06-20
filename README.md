# go-docker-skeleton

[![release](https://ghcr-badge.egpl.dev/bitxeno/go-docker-skeleton/latest_tag?label=docker%20latest)](https://github.com/bitxeno/go-docker-skeleton/pkgs/container/go-docker-skeleton)
[![image size](https://ghcr-badge.egpl.dev/bitxeno/go-docker-skeleton/size)](https://github.com/bitxeno/go-docker-skeleton/pkgs/container/go-docker-skeleton)
[![license](https://img.shields.io/github/license/bitxeno/go-docker-skeleton)](https://github.com/bitxeno/go-docker-skeleton/blob/master/LICENSE) 

Skeleton for run go service in docker


## Prerequisite

* Go 1.21+
* Node.js 16+
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

## Awesome Go Library

* [goutil](https://github.com/gookit/goutil): Helper Utils(600+)
* [lo](https://github.com/samber/lo): samber/lo is a Lodash-style Go library based on Go 1.18+ Generics.
* [validator](https://github.com/go-playground/validator): Package validator implements value validations for structs and individual fields based on tags.
* [gods](https://github.com/emirpasic/gods): GoDS (Go Data Structures) - Sets, Lists, Stacks, Maps, Trees, Queues, and much more
* [orderedmap](https://github.com/elliotchance/orderedmap): An ordered map in Go with amortized O(1) for Set, Get, Delete and Len.
* [cron](https://github.com/robfig/cron): a cron library for go
* [gocron](https://github.com/go-co-op/gocron): Easy and fluent Go cron scheduling.
* [gjson](https://github.com/tidwall/gjson): Get JSON values quickly - JSON parser for Go
* [ristretto](https://github.com/dgraph-io/ristretto): A high performance memory-bound Go cache
* [theine-go](https://github.com/Yiling-J/theine-go): High performance in-memory & hybrid cache with generics support
* [emitter](https://github.com/olebedev/emitter): Emits events in Go way, with wildcard, predicates, cancellation possibilities and many other good wins
* [event](https://github.com/gookit/event): Lightweight event manager and dispatcher implements by Go
* [asynq](https://github.com/hibiken/asynq): Simple, reliable, and efficient distributed task queue in Go
* [goja](https://github.com/dop251/goja): ECMAScript/JavaScript engine in pure Go
* [do](https://github.com/samber/do): ⚙️ A dependency injection toolkit based on Go 1.18+ Generics.
* [errors](https://github.com/go-errors/errors): errors with stacktraces for go
* [purego](https://github.com/ebitengine/purego): A library for calling C functions from Go without Cgo.

## Awesome Javascript Library

* [vue-sonner](https://github.com/xiaoluoboding/vue-sonner): 🔔 An opinionated toast component for Vue.
* [auto-animate](https://github.com/formkit/auto-animate): A zero-config, drop-in animation utility that adds smooth transitions to your web app.