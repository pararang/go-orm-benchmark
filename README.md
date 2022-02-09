# Go SQL Benchmark
## ORMs vs. SQl Builders vs. Raw
Several ORMs also support query builder, therefore the benchmark for these ORMs is made separately between those using the ORM-way and the query builder-way. 

To distinguish the benchmark of an ORM that uses the ORM-way or the builder-way, every initialization of the benchmark package name will be added the suffix -qb or -orm

### Environment

- go version go1.13 linux/amd64

### PostgreSQL

- PostgreSQL 12 for Linux on x86_64

### ORMs

All packages are the latest versions & run in no-cache mode.

- [gorm](https://github.com/go-gorm/gorm): orm
- [pg](https://github.com/go-pg/pg): orm, query builder
- [upper/db](https://github.com/upper/db): orm, query builder
- [sqlx](https://github.com/jmoiron/sqlx): orm
- [raw](https://pkg.go.dev/database/sql): raw/native go
- [ent](https://github.com/facebookincubator/ent): orm
- [goqu](https://github.com/doug-martin/goqu): query builder

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
                          ./+o+-       prrng@this
                  yyyyy- -yyyyyy+      OS: Ubuntu 20.04 focal
               ://+//////-yyyyyyo      Kernel: x86_64 Linux 5.4.0-99-generic
           .++ .:/++++++/-.+sss/`      Uptime: 21h 36m
         .:++o:  /++++++++/:--:/-      Packages: 2183
        o:+o+:++.`..```.-/oo+++++/     Shell: zsh 5.8
       .:+o:+o/.          `+sssoo+/    Resolution: 2732x768
  .++/+:+oo+o:`             /sssooo.   DE: GNOME 3.36.5
 /+++//+:`oo+o               /::--:.   WM: Mutter
 \+/+o+++`o++o               ++////.   WM Theme: Adwaita
  .++.o+++oo+:`             /dddhhh.   GTK Theme: Yaru [GTK2/3]
       .+.o+oo:.          `oddhhhh+    Icon Theme: Yaru
        \+.++o+o``-````.:ohdhhhhh+     Font: Ubuntu 10
         `:o+++ `ohhhhhhhhyo++os:      Disk: 66G / 124G (56%)
           .o:`.syhhhhhhh/.oo++o`      CPU: Intel Core i5-7200U @ 4x 3.1GHz [74.0°C]
               /osyyyyyyo++ooo+++/     GPU: NVIDIA GeForce 940MX
                   ````` +oo+++o\:     RAM: 6836MiB / 7825MiB
                          `oo++.      

➜  go-sql-benchmark git:(benchmark-emall) ✗ go run main.go -orm=all         
ent-orm
                   Insert:   2000     1.72s       862167 ns/op    5148 B/op    126 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000     2.30s      1149594 ns/op    5406 B/op    160 allocs/op
                     Read:   4000     1.94s       484983 ns/op    5344 B/op    147 allocs/op
      MultiRead limit 100:   2000     2.34s      1172208 ns/op   69345 B/op   2213 allocs/op
gopg-orm
                   Insert:   2000     0.90s       448271 ns/op    1075 B/op     12 allocs/op
      MultiInsert 100 row:    500     1.45s      2899561 ns/op   14747 B/op    215 allocs/op
                   Update:   2000     0.88s       438399 ns/op     904 B/op     13 allocs/op
                     Read:   4000     0.88s       219767 ns/op    1040 B/op     18 allocs/op
      MultiRead limit 100:   2000     0.92s       461323 ns/op   26397 B/op    624 allocs/op
gopg-qb
                   Insert:   2000     1.20s       601203 ns/op    1020 B/op     12 allocs/op
      MultiInsert 100 row:    500     1.36s      2724094 ns/op   14747 B/op    215 allocs/op
                   Update:   2000     0.82s       409445 ns/op     944 B/op     16 allocs/op
                     Read:   4000     0.75s       188200 ns/op    1096 B/op     16 allocs/op
      MultiRead limit 100:   2000     0.93s       462728 ns/op   26333 B/op    621 allocs/op
sqlx
                   Insert:   2000     0.68s       339856 ns/op     736 B/op     21 allocs/op
      MultiInsert 100 row:    500     1.51s      3028304 ns/op  142089 B/op   1423 allocs/op
                   Update:   2000     0.31s       157049 ns/op     744 B/op     21 allocs/op
                     Read:   4000     0.63s       157674 ns/op    1313 B/op     33 allocs/op
      MultiRead limit 100:   2000     0.92s       459749 ns/op   63977 B/op   1418 allocs/op
gorm
                   Insert:   2000     1.38s       691666 ns/op    6805 B/op    102 allocs/op
      MultiInsert 100 row:    500     1.39s      2788219 ns/op  197340 B/op   2703 allocs/op
                   Update:   2000     1.35s       677459 ns/op    5073 B/op     75 allocs/op
                     Read:   4000     0.76s       190807 ns/op    4633 B/op     96 allocs/op
      MultiRead limit 100:   2000     1.68s       842050 ns/op   61730 B/op   3740 allocs/op
upper-orm
                   Insert:   2000     5.51s      2754508 ns/op   36480 B/op   1624 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000     4.79s      2392746 ns/op   41508 B/op   1917 allocs/op
                     Read:   4000     0.91s       228207 ns/op    7844 B/op    345 allocs/op
      MultiRead limit 100:   2000     1.62s       808080 ns/op   56581 B/op   1884 allocs/op
raw
                   Insert:   2000     0.67s       337201 ns/op     736 B/op     21 allocs/op
      MultiInsert 100 row:    500     1.55s      3101677 ns/op  142089 B/op   1423 allocs/op
                   Update:   2000     0.30s       148623 ns/op     744 B/op     21 allocs/op
                     Read:   4000     0.59s       148663 ns/op     936 B/op     28 allocs/op
      MultiRead limit 100:   2000     0.77s       383283 ns/op   30256 B/op   1212 allocs/op
goqu-qb
                   Insert:   2000     0.87s       437377 ns/op    5884 B/op    223 allocs/op
      MultiInsert 100 row:    500     1.54s      3086399 ns/op  238404 B/op  11716 allocs/op
                   Update:   2000     0.77s       386562 ns/op    1927 B/op     43 allocs/op
                     Read:   4000     0.67s       168134 ns/op    2032 B/op     39 allocs/op
      MultiRead limit 100:   2000     0.63s       317048 ns/op    2448 B/op     45 allocs/op
upper-qb
                   Insert:   2000     1.69s       845211 ns/op   11390 B/op    594 allocs/op
      MultiInsert 100 row:    500     No support for multi insert
                   Update:   2000     1.62s       810380 ns/op    6245 B/op    330 allocs/op
                     Read:   4000     1.61s       402827 ns/op    7443 B/op    331 allocs/op
      MultiRead limit 100:   2000     1.42s       712065 ns/op   55022 B/op   2068 allocs/op

Reports: 

  2000 times - Insert
       raw:     0.67s       337201 ns/op     736 B/op     21 allocs/op
      sqlx:     0.68s       339856 ns/op     736 B/op     21 allocs/op
   goqu-qb:     0.87s       437377 ns/op    5884 B/op    223 allocs/op
  gopg-orm:     0.90s       448271 ns/op    1075 B/op     12 allocs/op
   gopg-qb:     1.20s       601203 ns/op    1020 B/op     12 allocs/op
      gorm:     1.38s       691666 ns/op    6805 B/op    102 allocs/op
  upper-qb:     1.69s       845211 ns/op   11390 B/op    594 allocs/op
   ent-orm:     1.72s       862167 ns/op    5148 B/op    126 allocs/op
 upper-orm:     5.51s      2754508 ns/op   36480 B/op   1624 allocs/op

   500 times - MultiInsert 100 row
   gopg-qb:     1.36s      2724094 ns/op   14747 B/op    215 allocs/op
      gorm:     1.39s      2788219 ns/op  197340 B/op   2703 allocs/op
  gopg-orm:     1.45s      2899561 ns/op   14747 B/op    215 allocs/op
      sqlx:     1.51s      3028304 ns/op  142089 B/op   1423 allocs/op
   goqu-qb:     1.54s      3086399 ns/op  238404 B/op  11716 allocs/op
       raw:     1.55s      3101677 ns/op  142089 B/op   1423 allocs/op
 upper-orm:     No support for multi insert
   ent-orm:     No support for multi insert
  upper-qb:     No support for multi insert

  2000 times - Update
       raw:     0.30s       148623 ns/op     744 B/op     21 allocs/op
      sqlx:     0.31s       157049 ns/op     744 B/op     21 allocs/op
   goqu-qb:     0.77s       386562 ns/op    1927 B/op     43 allocs/op
   gopg-qb:     0.82s       409445 ns/op     944 B/op     16 allocs/op
  gopg-orm:     0.88s       438399 ns/op     904 B/op     13 allocs/op
      gorm:     1.35s       677459 ns/op    5073 B/op     75 allocs/op
  upper-qb:     1.62s       810380 ns/op    6245 B/op    330 allocs/op
   ent-orm:     2.30s      1149594 ns/op    5406 B/op    160 allocs/op
 upper-orm:     4.79s      2392746 ns/op   41508 B/op   1917 allocs/op

  4000 times - Read
       raw:     0.59s       148663 ns/op     936 B/op     28 allocs/op
      sqlx:     0.63s       157674 ns/op    1313 B/op     33 allocs/op
   goqu-qb:     0.67s       168134 ns/op    2032 B/op     39 allocs/op
   gopg-qb:     0.75s       188200 ns/op    1096 B/op     16 allocs/op
      gorm:     0.76s       190807 ns/op    4633 B/op     96 allocs/op
  gopg-orm:     0.88s       219767 ns/op    1040 B/op     18 allocs/op
 upper-orm:     0.91s       228207 ns/op    7844 B/op    345 allocs/op
  upper-qb:     1.61s       402827 ns/op    7443 B/op    331 allocs/op
   ent-orm:     1.94s       484983 ns/op    5344 B/op    147 allocs/op

  2000 times - MultiRead limit 100
   goqu-qb:     0.63s       317048 ns/op    2448 B/op     45 allocs/op
       raw:     0.77s       383283 ns/op   30256 B/op   1212 allocs/op
      sqlx:     0.92s       459749 ns/op   63977 B/op   1418 allocs/op
  gopg-orm:     0.92s       461323 ns/op   26397 B/op    624 allocs/op
   gopg-qb:     0.93s       462728 ns/op   26333 B/op    621 allocs/op
  upper-qb:     1.42s       712065 ns/op   55022 B/op   2068 allocs/op
 upper-orm:     1.62s       808080 ns/op   56581 B/op   1884 allocs/op
      gorm:     1.68s       842050 ns/op   61730 B/op   3740 allocs/op
   ent-orm:     2.34s      1172208 ns/op   69345 B/op   2213 allocs/op

#################################

➜  go-sql-benchmark git:(benchmark-emall) ✗ go run main.go -orm=all -multi=10

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
