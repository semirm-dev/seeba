**Run**
```shell
docker-compose up
```

* Endpoint: localhost:8000 GET /music

**Tests**
```shell
go test ./... -v
```

* **etl:** main module, consisted of a few small interfaces which can be replaced as per new requirements (new featuers, breaking changes, new providers)
* **gateway:** exposes an api to get music gata
* **aoe:** provider specific implementations (_etl.Filter_ + _etl.Search_), new _etl.Importer_ implementation can also be provided (if needed)

> In order to add new provider or features with breaking changes:
> - create new package specific for the new provider and implement _etl.Import_, _etl.Filter_ and _etl.Search_ interfaces (either all 3 or only those with the breaking changes)
> - pass new provider implementation to _cmd/loader/main.go_