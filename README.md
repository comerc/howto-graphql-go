# graphql-go

https://www.howtographql.com/graphql-go/0-introduction/

## How to install MySQL

```bash
$ docker run -p 3306:3306 --name mysql -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -e MYSQL_DATABASE=hackernews -d mysql:latest
```

## How to enter to MySQL REPL

```bash
$ docker exec -it mysql mysql -u root
```

## How to Build `bin/migrate`

```bash
$ go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
```

## How to Migration

```bash
$ bin/migrate -database mysql://root:@/hackernews -path migrations up
```
