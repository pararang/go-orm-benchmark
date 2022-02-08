# Go SQL Benchmark
ORMs vs. SQl Builders vs. Raw

### Environment

- go version go1.13 linux/amd64

### PostgreSQL

- PostgreSQL 12 for Linux on x86_64

### ORMs

All packages are the latest versions & run in no-cache mode.

- [gorm](https://github.com/go-gorm/gorm)
- [pg](https://github.com/go-pg/pg)
- [upper/db](https://github.com/upper/db)
- [sqlx](https://github.com/jmoiron/sqlx)
- [raw](https://pkg.go.dev/database/sql)
- [ent](https://github.com/facebookincubator/ent)

### Run

```sh
# start database
docker run -d --rm \
  --name test-pg \
  -e POSTGRES_DB=test \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
    postgres:12-alpine -c log_statement=all

# all
go run main.go -multi=20 -orm=all

# portion
go run main.go -multi=20 -orm=gorm

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

```bash
go run main.go -orm=all -multi=10

raw
                   Insert:  20000     8.09s       404345 ns/op     736 B/op     21 allocs/op
      MultiInsert 100 row:   5000    16.23s      3246591 ns/op  141714 B/op   1421 allocs/op
                   Update:  20000     4.39s       219310 ns/op     744 B/op     21 allocs/op
                     Read:  40000     9.42s       235515 ns/op     936 B/op     28 allocs/op
      MultiRead limit 100:  20000     8.66s       433010 ns/op   30256 B/op   1212 allocs/op
sqlx
                   Insert:  20000     8.12s       406162 ns/op     736 B/op     21 allocs/op
      MultiInsert 100 row:   5000    17.10s      3420841 ns/op  141714 B/op   1421 allocs/op
                   Update:  20000     4.83s       241469 ns/op     744 B/op     21 allocs/op
                     Read:  40000     8.28s       206930 ns/op    1312 B/op     33 allocs/op
      MultiRead limit 100:  20000    10.37s       518380 ns/op   63976 B/op   1418 allocs/op
gopg-qb
                   Insert:  20000     9.67s       483438 ns/op     932 B/op     12 allocs/op
      MultiInsert 100 row:   5000    14.38s      2875075 ns/op   14798 B/op    215 allocs/op
                   Update:  20000     9.97s       498290 ns/op     944 B/op     16 allocs/op
                     Read:  40000    10.57s       264164 ns/op    1097 B/op     16 allocs/op
      MultiRead limit 100:  20000    10.74s       536987 ns/op   26281 B/op    621 allocs/op
upper-qb
                   Insert:  20000    14.10s       704859 ns/op   11386 B/op    594 allocs/op
      MultiInsert 100 row:   5000     No support for multi insert
                   Update:  20000    15.31s       765256 ns/op    6233 B/op    330 allocs/op
                     Read:  40000    22.39s       559864 ns/op    7440 B/op    331 allocs/op
      MultiRead limit 100:  20000    24.10s      1204861 ns/op   55017 B/op   2068 allocs/op
ent-orm
                   Insert:  20000    21.76s      1087768 ns/op    5131 B/op    126 allocs/op
      MultiInsert 100 row:   5000     No support for multi insert
                   Update:  20000    21.53s      1076285 ns/op    5400 B/op    160 allocs/op
                     Read:  40000    15.38s       384422 ns/op    5344 B/op    147 allocs/op
      MultiRead limit 100:  20000    14.02s       700785 ns/op   69344 B/op   2213 allocs/op
gopg-orm
                   Insert:  20000     7.63s       381683 ns/op     933 B/op     12 allocs/op
      MultiInsert 100 row:   5000    12.83s      2565195 ns/op   14746 B/op    215 allocs/op
                   Update:  20000     8.14s       406842 ns/op     904 B/op     13 allocs/op
                     Read:  40000     8.06s       201584 ns/op    1040 B/op     18 allocs/op
      MultiRead limit 100:  20000     8.36s       418242 ns/op   26374 B/op    624 allocs/op
goqu-qb
                   Insert:  20000     7.71s       385339 ns/op    5873 B/op    223 allocs/op
      MultiInsert 100 row:   5000    14.32s      2864088 ns/op  238404 B/op  11716 allocs/op
                   Update:  20000     7.07s       353306 ns/op    1927 B/op     43 allocs/op
                     Read:  40000     6.18s       154509 ns/op    2032 B/op     39 allocs/op
      MultiRead limit 100:  20000     5.82s       290772 ns/op    2448 B/op     45 allocs/op
gorm
                   Insert:  20000    17.01s       850704 ns/op    6755 B/op    102 allocs/op
      MultiInsert 100 row:   5000    12.29s      2458972 ns/op  197290 B/op   2703 allocs/op
                   Update:  20000    12.59s       629366 ns/op    5066 B/op     75 allocs/op
                     Read:  40000     7.03s       175727 ns/op    4632 B/op     96 allocs/op
      MultiRead limit 100:  20000    15.60s       780048 ns/op   61728 B/op   3740 allocs/op
upper-orm
                   Insert:  20000    43.24s      2162206 ns/op   36382 B/op   1622 allocs/op
      MultiInsert 100 row:   5000     No support for multi insert
                   Update:  20000    42.94s      2146835 ns/op   41484 B/op   1916 allocs/op
                     Read:  40000     8.07s       201722 ns/op    7840 B/op    345 allocs/op
      MultiRead limit 100:  20000    14.63s       731425 ns/op   56569 B/op   1884 allocs/op

Reports: 

 20000 times - Insert
  gopg-orm:     7.63s       381683 ns/op     933 B/op     12 allocs/op
   goqu-qb:     7.71s       385339 ns/op    5873 B/op    223 allocs/op
       raw:     8.09s       404345 ns/op     736 B/op     21 allocs/op
      sqlx:     8.12s       406162 ns/op     736 B/op     21 allocs/op
   gopg-qb:     9.67s       483438 ns/op     932 B/op     12 allocs/op
  upper-qb:    14.10s       704859 ns/op   11386 B/op    594 allocs/op
      gorm:    17.01s       850704 ns/op    6755 B/op    102 allocs/op
   ent-orm:    21.76s      1087768 ns/op    5131 B/op    126 allocs/op
 upper-orm:    43.24s      2162206 ns/op   36382 B/op   1622 allocs/op

  5000 times - MultiInsert 100 row
      gorm:    12.29s      2458972 ns/op  197290 B/op   2703 allocs/op
  gopg-orm:    12.83s      2565195 ns/op   14746 B/op    215 allocs/op
   goqu-qb:    14.32s      2864088 ns/op  238404 B/op  11716 allocs/op
   gopg-qb:    14.38s      2875075 ns/op   14798 B/op    215 allocs/op
       raw:    16.23s      3246591 ns/op  141714 B/op   1421 allocs/op
      sqlx:    17.10s      3420841 ns/op  141714 B/op   1421 allocs/op
  upper-qb:     No support for multi insert
   ent-orm:     No support for multi insert
 upper-orm:     No support for multi insert

 20000 times - Update
       raw:     4.39s       219310 ns/op     744 B/op     21 allocs/op
      sqlx:     4.83s       241469 ns/op     744 B/op     21 allocs/op
   goqu-qb:     7.07s       353306 ns/op    1927 B/op     43 allocs/op
  gopg-orm:     8.14s       406842 ns/op     904 B/op     13 allocs/op
   gopg-qb:     9.97s       498290 ns/op     944 B/op     16 allocs/op
      gorm:    12.59s       629366 ns/op    5066 B/op     75 allocs/op
  upper-qb:    15.31s       765256 ns/op    6233 B/op    330 allocs/op
   ent-orm:    21.53s      1076285 ns/op    5400 B/op    160 allocs/op
 upper-orm:    42.94s      2146835 ns/op   41484 B/op   1916 allocs/op

 40000 times - Read
   goqu-qb:     6.18s       154509 ns/op    2032 B/op     39 allocs/op
      gorm:     7.03s       175727 ns/op    4632 B/op     96 allocs/op
  gopg-orm:     8.06s       201584 ns/op    1040 B/op     18 allocs/op
 upper-orm:     8.07s       201722 ns/op    7840 B/op    345 allocs/op
      sqlx:     8.28s       206930 ns/op    1312 B/op     33 allocs/op
       raw:     9.42s       235515 ns/op     936 B/op     28 allocs/op
   gopg-qb:    10.57s       264164 ns/op    1097 B/op     16 allocs/op
   ent-orm:    15.38s       384422 ns/op    5344 B/op    147 allocs/op
  upper-qb:    22.39s       559864 ns/op    7440 B/op    331 allocs/op

 20000 times - MultiRead limit 100
   goqu-qb:     5.82s       290772 ns/op    2448 B/op     45 allocs/op
  gopg-orm:     8.36s       418242 ns/op   26374 B/op    624 allocs/op
       raw:     8.66s       433010 ns/op   30256 B/op   1212 allocs/op
      sqlx:    10.37s       518380 ns/op   63976 B/op   1418 allocs/op
   gopg-qb:    10.74s       536987 ns/op   26281 B/op    621 allocs/op
   ent-orm:    14.02s       700785 ns/op   69344 B/op   2213 allocs/op
 upper-orm:    14.63s       731425 ns/op   56569 B/op   1884 allocs/op
      gorm:    15.60s       780048 ns/op   61728 B/op   3740 allocs/op
  upper-qb:    24.10s      1204861 ns/op   55017 B/op   2068 allocs/op

```
