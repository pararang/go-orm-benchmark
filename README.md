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
gorm
                   Insert:   2000    46.24s     23120985 ns/op    6777 B/op    102 allocs/op
      MultiInsert 100 row:    500     8.14s     16278915 ns/op  195643 B/op   2702 allocs/op
                   Update:   2000    49.41s     24702651 ns/op    5064 B/op     75 allocs/op
                     Read:   4000     2.49s       621631 ns/op    4617 B/op     96 allocs/op
      MultiRead limit 100:   2000     4.86s      2431302 ns/op   60129 B/op   3740 allocs/op
pg
                   Insert:   2000    65.18s     32587960 ns/op    1075 B/op     12 allocs/op
      MultiInsert 100 row:    500     8.98s     17969529 ns/op   14745 B/op    215 allocs/op
                   Update:   2000    45.44s     22719287 ns/op     904 B/op     13 allocs/op
                     Read:   4000     1.74s       435649 ns/op    1024 B/op     18 allocs/op
      MultiRead limit 100:   2000     2.10s      1048390 ns/op   24794 B/op    624 allocs/op
upper
                   Insert:   2000    50.21s     25104760 ns/op   36447 B/op   1624 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000    52.33s     26164528 ns/op   41473 B/op   1917 allocs/op
                     Read:   4000     2.69s       672996 ns/op    7828 B/op    345 allocs/op
      MultiRead limit 100:   2000     5.87s      2935420 ns/op   54981 B/op   1884 allocs/op
raw
                   Insert:   2000    63.44s     31720677 ns/op     720 B/op     21 allocs/op
      MultiInsert 100 row:    500     8.05s     16105362 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     0.60s       300849 ns/op     728 B/op     21 allocs/op
                     Read:   4000     1.43s       356985 ns/op     920 B/op     28 allocs/op
      MultiRead limit 100:   2000     1.65s       825755 ns/op   28656 B/op   1212 allocs/op
sqlx
                   Insert:   2000    45.70s     22848368 ns/op     720 B/op     21 allocs/op
      MultiInsert 100 row:    500    12.02s     24037822 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     0.77s       386260 ns/op     728 B/op     21 allocs/op
                     Read:   4000     1.46s       363943 ns/op    1297 B/op     33 allocs/op
      MultiRead limit 100:   2000     2.11s      1056087 ns/op   62376 B/op   1418 allocs/op
ent
                   Insert:   2000    57.85s     28926652 ns/op    5140 B/op    126 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000    55.58s     27791732 ns/op    5399 B/op    160 allocs/op
                     Read:   4000     3.49s       873748 ns/op    5328 B/op    146 allocs/op
      MultiRead limit 100:   2000     3.84s      1920588 ns/op   69328 B/op   2212 allocs/op

Reports:

  2000 times - Insert
      sqlx:    45.70s     22848368 ns/op     720 B/op     21 allocs/op
      gorm:    46.24s     23120985 ns/op    6777 B/op    102 allocs/op
     upper:    50.21s     25104760 ns/op   36447 B/op   1624 allocs/op
       ent:    57.85s     28926652 ns/op    5140 B/op    126 allocs/op
       raw:    63.44s     31720677 ns/op     720 B/op     21 allocs/op
        pg:    65.18s     32587960 ns/op    1075 B/op     12 allocs/op

   500 times - MultiInsert 100 row
       raw:     8.05s     16105362 ns/op  140488 B/op   1423 allocs/op
      gorm:     8.14s     16278915 ns/op  195643 B/op   2702 allocs/op
        pg:     8.98s     17969529 ns/op   14745 B/op    215 allocs/op
      sqlx:    12.02s     24037822 ns/op  140488 B/op   1423 allocs/op
     upper:     No support for multi insert
       ent:     No support for multi insert

  2000 times - Update
       raw:     0.60s       300849 ns/op     728 B/op     21 allocs/op
      sqlx:     0.77s       386260 ns/op     728 B/op     21 allocs/op
        pg:    45.44s     22719287 ns/op     904 B/op     13 allocs/op
      gorm:    49.41s     24702651 ns/op    5064 B/op     75 allocs/op
     upper:    52.33s     26164528 ns/op   41473 B/op   1917 allocs/op
       ent:    55.58s     27791732 ns/op    5399 B/op    160 allocs/op

  4000 times - Read
       raw:     1.43s       356985 ns/op     920 B/op     28 allocs/op
      sqlx:     1.46s       363943 ns/op    1297 B/op     33 allocs/op
        pg:     1.74s       435649 ns/op    1024 B/op     18 allocs/op
      gorm:     2.49s       621631 ns/op    4617 B/op     96 allocs/op
     upper:     2.69s       672996 ns/op    7828 B/op    345 allocs/op
       ent:     3.49s       873748 ns/op    5328 B/op    146 allocs/op

  2000 times - MultiRead limit 100
       raw:     1.65s       825755 ns/op   28656 B/op   1212 allocs/op
        pg:     2.10s      1048390 ns/op   24794 B/op    624 allocs/op
      sqlx:     2.11s      1056087 ns/op   62376 B/op   1418 allocs/op
       ent:     3.84s      1920588 ns/op   69328 B/op   2212 allocs/op
      gorm:     4.86s      2431302 ns/op   60129 B/op   3740 allocs/op
     upper:     5.87s      2935420 ns/op   54981 B/op   1884 allocs/op
```
