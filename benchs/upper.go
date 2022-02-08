package benchs

import (
	"database/sql"
	"fmt"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var sess sqlbuilder.Database

func init() {
	st := NewSuite("upper-orm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, UpperInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, UpperInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, UpperUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, UpperRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, UpperReadSlice)

		d, _ := sql.Open("postgres", ORM_SOURCE)

		conn, err := postgresql.New(d)
		if err != nil {
			fmt.Println(err)
		}

		sess = conn
	}
}

func UpperInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		err := sess.Collection("models").InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperInsertMulti(b *B) {
	panic("No support for multi insert")
}

func UpperUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		err := sess.Collection("models").InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 20
		err := sess.Collection("models").UpdateReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		err := sess.Collection("models").InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		err := sess.Collection("models").Find().One(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			err := sess.Collection("models").InsertReturning(m)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := sess.Collection("models").Find("id > ?", 0).Limit(100).All(&models)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
