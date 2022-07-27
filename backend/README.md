## Infra
```shell
docker compose -f ./infra/docker-compose.yaml up -d
```

## Migrations

```shell
go run cmd/main m -up 0
go run cmd/main m -down 0
```

## API Server

```shell
go run cmd/main api
```
