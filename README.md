# gator

A multi-user command-line blog aggregation program.

## Pre-requisites

1. [PostgreSQL](https://www.postgresql.org/download/)
2. [Go](https://go.dev/doc/install)

Create a database in PostgreSQL and copy the connection string.

## Installing

```sh
go install github.com/donna-marijne/gator@latest
```

Create a file `~/.gatorconfig.json` containing your PostgreSQL connection string, e.g.:

```json
{
    "db_url": "postgres://localhost/gator?sslmode=disable"
}
```

## Usage

Run the scraping service with an interval of 5 minutes:

```sh
gator agg 5m
```

Register a user:

```sh
gator register alice
```

Log in as a user:

```sh
gator login alice
```

List the registered users:

```sh
gator users
```

Add a feed to be scraped:

```sh
gator addfeed "TechCrunch RSS Feed" https://techcrunch.com/feed/
```

List the added feeds:

```sh
gator feeds
```

Follow a feed:

```sh
gator follow https://techcrunch.com/feed/
```

Unfollow a feed:

```sh
gator unfollow https://techcrunch.com/feed/
```

List the followed feeds:

```sh
gator following
```

Show the most recent posts across all followed feeds:

```sh
gator browse 10
```
