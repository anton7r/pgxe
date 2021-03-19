# PGXE - PGX Extension

[![](https://godoc.org/github.com/anton7r/pgxe?status.svg)](https://pkg.go.dev/github.com/anton7r/pgxe)
![Coverage](./coverage_badge.png)

Extends from the jackc/pgx library.

Utilizes the pgx library under the hood to achieve fast performance.

It also uses a modified fork of scany (`github.com/anton7r/pgx-scany` which removes unneccessary imports) to scan database rows into structs.

## Installation

`go get -u github.com/anton7r/pgxe`

## Features

- Simplified SQLX like API
- Named Queries Supported and they are fast

## Code example

The code example is not guranteed to have correct syntax, so if you spot a mistake please file an issue :)

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
    })

    query := "SELECT * FROM products WHERE id = :id"
    rows, err := db.NamedQuery(query, &Product{Id:4})
    if err != nil {
        println(err.Error())
        return
    }
}
```
