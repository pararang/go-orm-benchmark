package benchs

import (
	"fmt"

	// "database/sql"

	pg "github.com/go-pg/pg/v10"
)

var pgdbBuilder *pg.DB

func init() {
	st := NewSuite("gopg-qb")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, PgBuilderInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, PgBuilderInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, PgBuilderUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, PgBuilderRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, PgBuilderReadSlice)

		opts, _ := pg.ParseURL(ORM_SOURCE)
		pgdbBuilder = pg.Connect(opts)
	}
}

func PgBuilderInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.ID = 0
		_, err := pgdbBuilder.Model(m).Insert()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgBuilderInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}

		_, err := pgdbBuilder.Model(&ms).Insert()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgBuilderUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdbBuilder.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 20
		_, err := pgdbBuilder.Model(m).Column("age").WherePK().Update()
		if err != nil {

			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgBuilderRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := pgdbBuilder.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		mRes := new(Model)
		err := pgdbBuilder.Model(mRes).Limit(1).Select()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgBuilderReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			m.ID = 0
			if err := pgdbBuilder.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		err := pgdbBuilder.Model(&models).Limit(100).Select()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
