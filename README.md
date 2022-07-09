**Run**
```shell
docker-compose up
```

* Endpoint: localhost:8000 GET /music

**Tests**
```shell
go test ./... -v
```

* etl: main module
* aoe: provider specific implementations (import + filter + search api)

> In order to add new provider or features with breaking changes:
> - create new package specific for the new provider and implement _etl.Import_, _etl.Filter_ and _etl.Search_ interfaces (either all 3 or only those with the breaking changes)
> - pass new provider implementation to _cmd/gateway/main.go_