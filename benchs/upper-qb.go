package benchs

import (
	"database/sql"
	"fmt"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
)

var sessqb sqlbuilder.Database

func init() {
	st := NewSuite("upper-qb")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, UpperQueryBuilderInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, UpperQueryBuilderInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, UpperQueryBuilderUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, UpperQueryBuilderRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, UpperQueryBuilderReadSlice)

		d, _ := sql.Open("postgres", ORM_SOURCE)

		conn, err := postgresql.New(d)
		if err != nil {
			fmt.Println(err)
		}

		sessqb = conn
	}
}

func UpperQueryBuilderInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		query := sessqb.InsertInto("models").Values(m)
		_, err := query.Exec()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperQueryBuilderInsertMulti(b *B) {
	panic("No support for multi insert")
}

func UpperQueryBuilderUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		err := sessqb.Collection("models").InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 20
		query := sessqb.Update("models").Set("age", i).Where("id = ?", m.ID)
		_, err := query.Exec()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperQueryBuilderRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		err := sessqb.Collection("models").InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		query := sessqb.Select("*").From("models").Where("id = ?", m.ID)
		err := query.One(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperQueryBuilderReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			err := sessqb.Collection("models").InsertReturning(m)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		query := sessqb.Select("*").From("models").Limit(100)
		err := query.All(&models)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
