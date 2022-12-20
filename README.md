## Using migrations

```shell
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
```

```shell
migrate create -ext sql -dir database/migrations -seq create_users_table
```