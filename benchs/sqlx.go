package benchs

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var sqlxdb *sqlx.DB

func init() {
	st := NewSuite("sqlx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, SqlxInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, SqlxInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, SqlxUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, SqlxRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, SqlxReadSlice)

		db, err := sqlx.Connect("postgres", ORM_SOURCE)
		checkErr(err)
		sqlxdb = db
	}
}

func SqlxInsert(b *B) {
	var m *Model
	var stmt *sqlx.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		stmt, err = sqlxdb.Preparex(rawInsertSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		// pq does not support the LastInsertId method.
		_, err := stmt.Exec(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
	})

	var valuesSQL string
	counter := 1
	for i := 0; i < 100; i++ {
		hoge := ""
		for j := 0; j < 7; j++ {
			if j != 6 {
				hoge += "$" + strconv.Itoa(counter) + ","
			} else {
				hoge += "$" + strconv.Itoa(counter)
			}
			counter++

		}
		if i != 99 {
			valuesSQL += "(" + hoge + "),"
		} else {
			valuesSQL += "(" + hoge + ")"
		}
	}

	for i := 0; i < b.N; i++ {
		nFields := 7
		query := rawInsertBaseSQL + valuesSQL
		args := make([]interface{}, len(ms)*nFields)
		for j := range ms {
			offset := j * nFields
			args[offset+0] = ms[j].Name
			args[offset+1] = ms[j].Title
			args[offset+2] = ms[j].Fax
			args[offset+3] = ms[j].Web
			args[offset+4] = ms[j].Age
			args[offset+5] = ms[j].Right
			args[offset+6] = ms[j].Counter
		}
		// pq does not support the LastInsertId method.
		_, err := sqlxdb.Exec(query, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxUpdate(b *B) {
	var m *Model
	var stmt *sqlx.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		sqlxdb.MustExec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		stmt, err = sqlxdb.Preparex(rawUpdateSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		_, err := stmt.Exec(m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.ID)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxRead(b *B) {
	var m *Model
	var stmt *sqlx.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		sqlxdb.MustExec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		stmt, err = sqlxdb.Preparex(rawSelectSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		var mout Model
		if err := stmt.Get(&mout, 1); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxReadSlice(b *B) {
	var m *Model
	var stmt *sqlx.Stmt
	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			sqlxdb.MustExec(rawInsertSQL, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		}
		stmt, err = sqlxdb.Preparex(rawSelectMultiSQL)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	defer stmt.Close()

	for i := 0; i < b.N; i++ {
		models := make([]Model, 100)
		if err := stmt.Select(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
