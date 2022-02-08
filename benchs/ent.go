package benchs

import (
	"context"
	"fmt"

	"github.com/dre1080/go-orm-benchmark/benchs/ent"
)

var entdb *ent.Client

func init() {
	st := NewSuite("ent-orm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, EntInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, EntInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, EntUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, EntRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, EntReadSlice)

		conn, err := ent.Open("postgres", ORM_SOURCE)
		if err != nil {
			fmt.Println(err)
		}
		entdb = conn
	}
}

func EntCreate() (*ent.Model, error) {
	return entdb.Model.
		Create().
		SetName("Orm Benchmark").
		SetTitle("Just a Benchmark for fun").
		SetFax("99909990").SetWeb("http://blog.milkpod29.me").
		SetAge(100).
		SetRight(true).
		SetCounter(1000).
		Save(context.Background())
}

func EntInsert(b *B) {
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		_, err := EntCreate()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func EntInsertMulti(b *B) {
	panic("No support for multi insert")
}

func EntUpdate(b *B) {
	var m *ent.Model
	wrapExecute(b, func() {
		var err error
		initDB()
		m, err = EntCreate()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := m.Update().SetAge(20).Save(context.Background())
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func EntRead(b *B) {
	var m *ent.Model
	wrapExecute(b, func() {
		var err error
		initDB()
		m, err = EntCreate()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := entdb.Model.
			Get(context.Background(), m.ID)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func EntReadSlice(b *B) {
	wrapExecute(b, func() {
		initDB()
		for i := 0; i < 100; i++ {
			_, err := EntCreate()
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := entdb.Model.Query().Limit(100).All(context.Background())
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
