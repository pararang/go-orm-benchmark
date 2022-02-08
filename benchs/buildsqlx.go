package benchs

import (
	"fmt"
	"log"

	"github.com/arthurkushman/buildsqlx"
	// _ "github.com/lib/pq"
)

var buildsqlxdb *buildsqlx.DB

func init() {
	return

	//TODO: gagal insert jika ada boolean, gagal binding

	st := NewSuite("buildsqlx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, BuildsqlxInsert)
		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, SqlxInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, SqlxUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, SqlxRead)
		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, SqlxReadSlice)

		buildsqlxCon := buildsqlx.NewConnection("postgres", ORM_SOURCE)
		buildsqlxdb = buildsqlx.NewDb(buildsqlxCon)
	}
}

func BuildsqlxInsert(b *B) {
	var m *Farmer
	var newRow map[string]interface{}
	wrapExecute(b, func() {
		initDB()
		m = NewFarmer()
		newRow = m.MapStringInterface()
	})

	for i := 0; i < b.N; i++ {
		log.Println("OK++++++++++++++++++++++++++++++++++++++++++++1")
		log.Printf("%#v", newRow)
		err := buildsqlxdb.Table("farmers").Insert(newRow)
		log.Println("OK++++++++++++++++++++++++++++++++++++++++++++2")
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BuildsqlxInsertMulti(b *B) {
	var newRows []map[string]interface{}
	wrapExecute(b, func() {
		initDB()
		newRows = make([]map[string]interface{}, 0, 100)
		for i := 0; i < 100; i++ {
			newRows = append(newRows, NewFarmer().MapStringInterface())
		}
	})

	for i := 0; i < b.N; i++ {
		err := buildsqlxdb.Table("farmers").InsertBatch(newRows)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BuildsqlxUpdate(b *B) {
	var (
		newRow map[string]interface{}
		m      *Farmer
		newId  uint64
	)

	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewFarmer()
		newRow = m.MapStringInterface()
		newId, err = buildsqlxdb.Table("farmers").InsertGetId(newRow)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		m.Age = 500
		_, err := buildsqlxdb.Table("farmers").Where("id", "=", newId).Update(m.MapStringInterface())
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BuildsqlxRead(b *B) {
	var (
		newRow map[string]interface{}
		m      *Farmer
	)

	wrapExecute(b, func() {
		var err error
		initDB()
		m = NewFarmer()
		newRow = m.MapStringInterface()
		err = buildsqlxdb.Table("farmers").Insert(newRow)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := buildsqlxdb.Table("farmers").Select("*").Limit(1).Get()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func BuildsqlxReadSlice(b *B) {
	var (
		newRow map[string]interface{}
		m      *Farmer
		err    error
	)
	wrapExecute(b, func() {
		initDB()
		m = NewFarmer()
		newRow = m.MapStringInterface()
		for i := 0; i < 100; i++ {
			err = buildsqlxdb.Table("farmers").Insert(newRow)
			if err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := buildsqlxdb.Table("farmers").Select("*").Limit(100).Get()
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
