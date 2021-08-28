# PGXE - PGX Extension

[![Refrence](https://godoc.org/github.com/anton7r/pgxe?status.svg)](https://pkg.go.dev/github.com/anton7r/pgxe)
![Coverage](./coverage_badge.png)

This library aims to reduce the cognitive load while trying to make sql queries and hastens development.

Query preparation is fast (for most queries 500ns - 800ns) because it cuts certain corners such as not handling comments at all, which may be problematic for you or not.
It is recommended that in production you use just pgx because the overhead of this library might be sometimes a bit too much in systems that need to scale extremely well

Extends from the jackc/pgx library.

Utilizes the pgx library under the hood to achieve fast performance.

It also uses a modified fork of scany (`github.com/anton7r/pgx-scany` which removes unneccessary imports) to scan database rows into structs.

## Installation

`go get -u github.com/anton7r/pgxe`

## Features

- Simplified SQLX like API
- Named Queries Supported and they are fast

## Code example

The code example is not guranteed to have correct go-syntax, so if you spot a mistake please file an issue :)

```go
import "github.com/anton7r/pgxe"

type Product struct {
    Id int
    Name string
    Price decimal??
}

func main() {
    db := pgxe.Connect(pgxe.Connection{
        User:     "admin",
        Password: "superSecretPassword",
        DbName:   "postgres",
        DbPort:   "5432",

        Logger:   logger //pgx.Logger
    })

    //The id placeholder in the query has to be same case as defined in the struct so that the field can be found
    //Case insensitivity would decrease the performance of our lexer by at the worst case ~5x, so it is better for it to be case sensitive
    query := "SELECT * FROM products WHERE id = :Id"
    rows, err := db.NamedQuery(query, &Product{Id:4})
    if err != nil {
        println(err.Error())
        return
    }
}
```
