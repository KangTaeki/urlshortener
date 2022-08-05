# urlshortener

[![Actions Status](https://github.com/banksalad/urlshortener/workflows/ci/badge.svg)](https://github.com/banksalad/urlshortener/actions) ![Golang Badge](https://badgen.net/badge/Language/Go/cyan) ![GRPC Badge](https://badgen.net/badge/Use/gRPC/blue)

> urlshortener

## Docs
<!-- TODO: Update to the actual document url -->
- [TechSpec](https://docs.google.com/document/d/your_tech_spec_path)
- [Go in Banksalad](https://docs.google.com/document/d/1mPUGKlfA6pFLMUuUCHv54ejnUDrrldJ5z06AbvinRQA)

## Getting Started

### Start a Server
```sh
$ git config --global url."https://${GITHUB_ACCESS_TOKEN}@github.com/".insteadOf "https://github.com/"  # insert your github access token
$ make init
$ make run

# Use Docker
$ docker build --build-arg GH_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN} --tag urlshortener .  # insert your github access token
$ docker run --rm -p 18081:18081 -p 18082:18082 urlshortener

# Use Onebox
$ make deploy-to-local-k8s
```

### Test & Lint
```sh
$ make test

$ make lint
```

## APIs
<!-- TODO: Update to actual urls -->
- [urlshortener.proto](https://github.com/banksalad/idl/blob/master/protos/apis/v1/urlshortener/urlshortener.proto)
- [urlshortener.swagger.json](https://github.com/banksalad/idl/blob/master/gen/swagger/apis/v1/urlshortener/urlshortener.swagger.json)

## Directory Structure
```
.
├── client.go # dependency service 들에 대한 client
├── cmd       # server를 실행시키기 위한 main.go
│   └── ...
├── config    # 설정 파일
│   └── ...
└── server
    ├── grpc_server.go  # main gRPC server
    ├── http_server.go  # HTTP <-> gRPC 변환해주는 grpc-gateway layer
    └── handler         # gRPC server handlers
```
