package benchs

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var dbGoqu *goqu.Database

func init() {
	st := NewSuite("goqu-qb")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, GoquInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, GoquInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, GoquUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, GoquRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, GoquReadSlice)

		dbCon, err := sql.Open("postgres", ORM_SOURCE)
		if err != nil {
			fmt.Println(err)
			return
		}

		dbGoqu = goqu.New("postgres", dbCon)
	}
}

func GoquInsert(b *B) {
	var (
		m *Model
	)
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		gqData := dbGoqu.Insert("models").Rows(m)

		insertSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(insertSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GoquInsertMulti(b *B) {
	var ms []Model
	wrapExecute(b, func() {
		initDB()

		ms = make([]Model, 0, 100)
		for i := 0; i < 100; i++ {
			newModel := NewModel()
			ms = append(ms, *newModel)
		}
	})

	for i := 0; i < b.N; i++ {
		gqData := dbGoqu.Insert("models").Rows(ms)
		insertSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(insertSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GoquUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gqData := dbGoqu.Insert("models").Rows(m)
		insertSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(insertSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = int(i)
		gqData := dbGoqu.Update("models").Set(goqu.Record{"age": m.Age})
		updateSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(updateSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GoquRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gqData := dbGoqu.Insert("models").Rows(m)
		insertSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(insertSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		gqData := dbGoqu.From("models").Select("*")
		selectSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(selectSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GoquReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		gqData := dbGoqu.Insert("models").Rows(m)
		insertSQL, args, _ := gqData.ToSQL()
		for i := 0; i < 100; i++ {
			_, err := dbGoqu.Exec(insertSQL, args...)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		gqData := dbGoqu.From("models").Select("*").Limit(100)
		selectSQL, args, _ := gqData.ToSQL()
		_, err := dbGoqu.Exec(selectSQL, args...)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
