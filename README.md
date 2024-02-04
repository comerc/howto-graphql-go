
How to install MySQL

```bash
$ docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=hackernews -d mysql:latest
```

How to enter to MySQL REPL

```bash
$ docker exec -it mysql bash
$ mysql -u root -p
```

How to Build `bin/migrate`

```bash
$ go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
```

How to Migration

```bash
$ bin/migrate -database mysql://root:dbpass@/hackernews -path migrations up
```

