# oauth2-server

`tamanyan/oauth2-server` is a implementation of an [OAuth 2.0](https://tools.ietf.org/html/rfc6749) authorization server written in GO.

# How to build

```sh
make engine
```

# How to run

```sh
./bin/engine
```

# How to add module into app dir

## 1. Download andersjanmyr/goose

Download goose in order to generate template files.

[andersjanmyr/goose](https://github.com/andersjanmyr/goose)

```
# OS X
$ curl -L https://github.com/andersjanmyr/goose/releases/download/v1.5.0/goose-osx > goose

# Linux
$ curl -L https://github.com/andersjanmyr/goose/releases/download/v1.5.0/goose-linux > goose

# Make executable
$ chmod a+x ./goose
```

For example, let's create `article` module

```sh
> MODULE=article make gen-template
./goose --verbose -outputdir app/article -templatedir ./.template/ -data 'name=article'  .  article
OPTIONS:
verbose: true
force: false
interactive: false
template: .
name: article
templateDir: ./.template/
outputDir: app/article
data: map[DATA:map[name:article] NAME:article]
Generating file app/article/http/controller/article_controller.go
Generating file app/article/http/request/article_request.go
Generating file app/article/http/response/article_response.go
Generating file app/article/repository/article_repository.go
Generating file app/article/repository.go
Generating file app/article/usecase/article_ucase.go
Generating file app/article/usecase.go
```

# Acknowledgements

`tamanyan/oauth2-server` uses [go-oauth2/oauth2](https://github.com/go-oauth2/oauth2).

```
MIT License

Copyright (c) 2016 Lyric
```
