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
sqlx
                   Insert:   2000    38.48s     19241564 ns/op     720 B/op     21 allocs/op
      MultiInsert 100 row:    500     6.64s     13279301 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     0.66s       329501 ns/op     728 B/op     21 allocs/op
                     Read:   4000     1.44s       359778 ns/op    1297 B/op     33 allocs/op
      MultiRead limit 100:   2000     2.44s      1217624 ns/op   62376 B/op   1418 allocs/op
raw
                   Insert:   2000    55.04s     27521602 ns/op     720 B/op     21 allocs/op
      MultiInsert 100 row:    500    11.33s     22650811 ns/op  140488 B/op   1423 allocs/op
                   Update:   2000     0.66s       329317 ns/op     728 B/op     21 allocs/op
                     Read:   4000     1.04s       260353 ns/op     920 B/op     28 allocs/op
      MultiRead limit 100:   2000     2.02s      1009911 ns/op   28656 B/op   1212 allocs/op
pg
                   Insert:   2000    36.17s     18085514 ns/op    1042 B/op     12 allocs/op
      MultiInsert 100 row:    500    14.56s     29127838 ns/op   14745 B/op    215 allocs/op
                   Update:   2000    48.01s     24002962 ns/op     904 B/op     13 allocs/op
                     Read:   4000     2.83s       706404 ns/op    1024 B/op     18 allocs/op
      MultiRead limit 100:   2000     2.94s      1468444 ns/op   24762 B/op    624 allocs/op
upper
                   Insert:   2000    43.78s     21891197 ns/op   36454 B/op   1624 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000    44.97s     22485917 ns/op   41480 B/op   1918 allocs/op
                     Read:   4000     3.02s       754957 ns/op    7828 B/op    345 allocs/op
      MultiRead limit 100:   2000     4.32s      2159214 ns/op   54981 B/op   1884 allocs/op
gorm
                   Insert:   2000    35.59s     17797016 ns/op    6096 B/op     96 allocs/op
      MultiInsert 100 row:    500     7.61s     15212086 ns/op  194938 B/op   2696 allocs/op
                   Update:   2000    40.21s     20102681 ns/op    6716 B/op     97 allocs/op
                     Read:   4000     1.83s       457193 ns/op    3408 B/op     71 allocs/op
      MultiRead limit 100:   2000     4.93s      2464402 ns/op   60520 B/op   3752 allocs/op

Reports:

  2000 times - Insert
      gorm:    35.59s     17797016 ns/op    6096 B/op     96 allocs/op
        pg:    36.17s     18085514 ns/op    1042 B/op     12 allocs/op
      sqlx:    38.48s     19241564 ns/op     720 B/op     21 allocs/op
     upper:    43.78s     21891197 ns/op   36454 B/op   1624 allocs/op
       raw:    55.04s     27521602 ns/op     720 B/op     21 allocs/op

   500 times - MultiInsert 100 row
      sqlx:     6.64s     13279301 ns/op  140488 B/op   1423 allocs/op
      gorm:     7.61s     15212086 ns/op  194938 B/op   2696 allocs/op
       raw:    11.33s     22650811 ns/op  140488 B/op   1423 allocs/op
        pg:    14.56s     29127838 ns/op   14745 B/op    215 allocs/op
     upper:     No support for multi insert

  2000 times - Update
       raw:     0.66s       329317 ns/op     728 B/op     21 allocs/op
      sqlx:     0.66s       329501 ns/op     728 B/op     21 allocs/op
      gorm:    40.21s     20102681 ns/op    6716 B/op     97 allocs/op
     upper:    44.97s     22485917 ns/op   41480 B/op   1918 allocs/op
        pg:    48.01s     24002962 ns/op     904 B/op     13 allocs/op

  4000 times - Read
       raw:     1.04s       260353 ns/op     920 B/op     28 allocs/op
      sqlx:     1.44s       359778 ns/op    1297 B/op     33 allocs/op
      gorm:     1.83s       457193 ns/op    3408 B/op     71 allocs/op
        pg:     2.83s       706404 ns/op    1024 B/op     18 allocs/op
     upper:     3.02s       754957 ns/op    7828 B/op    345 allocs/op

  2000 times - MultiRead limit 100
       raw:     2.02s      1009911 ns/op   28656 B/op   1212 allocs/op
      sqlx:     2.44s      1217624 ns/op   62376 B/op   1418 allocs/op
        pg:     2.94s      1468444 ns/op   24762 B/op    624 allocs/op
     upper:     4.32s      2159214 ns/op   54981 B/op   1884 allocs/op
      gorm:     4.93s      2464402 ns/op   60520 B/op   3752 allocs/op
```
