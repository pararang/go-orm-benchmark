package benchs

import (
	"database/sql"
	"fmt"
	"os"
)

type Model struct {
	ID      int    `db:"id,omitempty" gorm:"primary_key" pg:",pk" goqu:"skipinsert"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "ORM x SQL Builder x Native"
	m.Title = "eMall BE Standard"
	m.Fax = "99909990"
	m.Web = "https://efishery.com"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

type Farmer struct {
	ID      int    `db:"id,omitempty" gorm:"primary_key" pg:",pk"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	IsRight bool   `db:"is_right"`
	Counter int64  `db:"counter"`
}

func NewFarmer() *Farmer {
	m := new(Farmer)
	m.Name = "Bench"
	m.Title = "GoPostgres"
	m.Fax = "99909990"
	m.Web = "https://pararang.com"
	m.Age = 100
	m.IsRight = true
	m.Counter = 1000

	return m
}

func (m *Farmer) MapStringInterface() map[string]interface{} {
	return map[string]interface{}{
		// "id":       m.ID,
		"name":     m.Name,
		"title":    m.Title,
		"fax":      m.Fax,
		"web":      m.Web,
		"age":      m.Age,
		"is_right": m.IsRight,
		"counter":  m.Counter,
	}
}

var (
	ORM_MULTI    int
	ORM_MAX_IDLE int
	ORM_MAX_CONN int
	ORM_SOURCE   string
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
}

func initDB() {

	sqls := []string{
		`DROP TABLE IF EXISTS models;`,
		`CREATE TABLE models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,

		`DROP TABLE IF EXISTS farmers;`,
		`CREATE TABLE farmers (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			is_right boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT farmers_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", ORM_SOURCE)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}
}
