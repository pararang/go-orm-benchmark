package benchs

import (
	"fmt"

	pg "github.com/go-pg/pg/v10"
)

var pgdbORM *pg.DB

func init() {
	st := NewSuite("gopg-orm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, PgORMInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, PgORMInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, PgORMUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, PgORMRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, PgORMReadSlice)

		opts, _ := pg.ParseURL(ORM_SOURCE)
		pgdbORM = pg.Connect(opts)
	}
}

func PgORMInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		if err := pgdbORM.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgORMInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()		
	})

	for i := 0; i < b.N; i++ {
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}

		if err := pgdbORM.Insert(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgORMUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdbORM.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 20
		if err := pgdbORM.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgORMRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdbORM.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pgdbORM.Select(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgORMReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			if err := pgdbORM.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := pgdbORM.Model(&models).Where("id > ?", 0).Limit(100).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
