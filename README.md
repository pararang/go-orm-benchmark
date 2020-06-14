# Go ORM Benchmark

### Environment

- go version go1.14.2 linux/amd64

### PostgreSQL

- PostgreSQL 12 for Linux on x86_64

### ORMs

All packages are the latest versions & run in no-cache mode.

- [gorm](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [upper/db](https://github.com/upper/db)
- [sqlx](https://github.com/jmoiron/sqlx)
- [raw](https://pkg.go.dev/database/sql)

### Run

```sh
# start database
docker run -d --rm \
  --name test-pg \
  -e POSTGRES_DB=test \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
    postgres:12-alpine

# all
go run github.com/dre1080/go-orm-benchmark -multi=20 -orm=all

# portion
go run github.com/dre1080/go-orm-benchmark -multi=20 -orm=gorm

# stop database
docker stop test-pg
```

### Example Reports

From the left:

- ORM name
- Total execution time (less is better)
- Time to execute once (less is better)
- Memory size allocated for each execution (smaller is better)
- Number of memory allocations (memory allocation/allocation) performed in one execution (smaller is better)

```
pg
                   Insert:   2000    32.18s     16091028 ns/op    1042 B/op     12 allocs/op
      MultiInsert 100 row:    500     9.81s     19621318 ns/op   14749 B/op    215 allocs/op
                   Update:   2000    34.30s     17151712 ns/op     904 B/op     13 allocs/op
                     Read:   4000     1.92s       478763 ns/op    1024 B/op     18 allocs/op
      MultiRead limit 100:   2000     2.22s      1110607 ns/op   24761 B/op    624 allocs/op
raw
                   Insert:   2000    35.68s     17840277 ns/op     720 B/op     21 allocs/op
      MultiInsert 100 row:    500     6.31s     12620714 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     0.65s       324239 ns/op     728 B/op     21 allocs/op
                     Read:   4000     1.18s       293964 ns/op     920 B/op     28 allocs/op
      MultiRead limit 100:   2000     1.75s       876910 ns/op   28656 B/op   1212 allocs/op
upper
                   Insert:   2000    34.96s     17478780 ns/op   36452 B/op   1624 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000    45.37s     22685399 ns/op   41480 B/op   1918 allocs/op
                     Read:   4000     2.08s       520780 ns/op    7828 B/op    345 allocs/op
      MultiRead limit 100:   2000     4.47s      2237225 ns/op   54981 B/op   1884 allocs/op
gorm
                   Insert:   2000    46.01s     23002695 ns/op    5807 B/op     93 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000    37.55s     18775701 ns/op    6671 B/op     95 allocs/op
                     Read:   4000     3.68s       919299 ns/op    4000 B/op     84 allocs/op
      MultiRead limit 100:   2000     5.09s      2544694 ns/op   61104 B/op   3765 allocs/op
sqlx
                   Insert:   2000    41.16s     20581732 ns/op     896 B/op     22 allocs/op
      MultiInsert 100 row:    500     7.56s     15126083 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     1.33s       664444 ns/op     904 B/op     22 allocs/op
                     Read:   4000     1.40s       351112 ns/op    1601 B/op     39 allocs/op
      MultiRead limit 100:   2000     2.78s      1390878 ns/op   58640 B/op   1724 allocs/op

Reports:

  2000 times - Insert
        pg:    32.18s     16091028 ns/op    1042 B/op     12 allocs/op
     upper:    34.96s     17478780 ns/op   36452 B/op   1624 allocs/op
       raw:    35.68s     17840277 ns/op     720 B/op     21 allocs/op
      sqlx:    41.16s     20581732 ns/op     896 B/op     22 allocs/op
      gorm:    46.01s     23002695 ns/op    5807 B/op     93 allocs/op

   500 times - MultiInsert 100 row
       raw:     6.31s     12620714 ns/op  140488 B/op   1423 allocs/op
      sqlx:     7.56s     15126083 ns/op  140488 B/op   1423 allocs/op
        pg:     9.81s     19621318 ns/op   14749 B/op    215 allocs/op
     upper:     No support for multi insert
      gorm:     No support for multi insert

  2000 times - Update
       raw:     0.65s       324239 ns/op     728 B/op     21 allocs/op
      sqlx:     1.33s       664444 ns/op     904 B/op     22 allocs/op
        pg:    34.30s     17151712 ns/op     904 B/op     13 allocs/op
      gorm:    37.55s     18775701 ns/op    6671 B/op     95 allocs/op
     upper:    45.37s     22685399 ns/op   41480 B/op   1918 allocs/op

  4000 times - Read
       raw:     1.18s       293964 ns/op     920 B/op     28 allocs/op
      sqlx:     1.40s       351112 ns/op    1601 B/op     39 allocs/op
        pg:     1.92s       478763 ns/op    1024 B/op     18 allocs/op
     upper:     2.08s       520780 ns/op    7828 B/op    345 allocs/op
      gorm:     3.68s       919299 ns/op    4000 B/op     84 allocs/op

  2000 times - MultiRead limit 100
       raw:     1.75s       876910 ns/op   28656 B/op   1212 allocs/op
        pg:     2.22s      1110607 ns/op   24761 B/op    624 allocs/op
      sqlx:     2.78s      1390878 ns/op   58640 B/op   1724 allocs/op
     upper:     4.47s      2237225 ns/op   54981 B/op   1884 allocs/op
      gorm:     5.09s      2544694 ns/op   61104 B/op   3765 allocs/op
```
