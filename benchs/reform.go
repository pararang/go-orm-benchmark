package benchs

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	"gopkg.in/reform.v1"
// 	"gopkg.in/reform.v1/dialects/postgresql"
// )

// var reformdb *reform.DB

// perlu banyak tewak krn method2 nya hanya menerima interface khusus dr reform

// func init() {
// 	st := NewSuite("gorm")
// 	st.InitF = func() {
// 		st.AddBenchmark("Insert", 2000*ORM_MULTI, ReformInsert)
// 		st.AddBenchmark("MultiInsert 100 row", 500*ORM_MULTI, ReformInsertMulti)
// 		st.AddBenchmark("Update", 2000*ORM_MULTI, ReformUpdate)
// 		st.AddBenchmark("Read", 4000*ORM_MULTI, ReformRead)
// 		st.AddBenchmark("MultiRead limit 100", 2000*ORM_MULTI, ReformReadSlice)

// 		sqlDB, err := sql.Open("postgres", "postgres://127.0.0.1:5432/database")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer sqlDB.Close()

// 		logger := log.New(os.Stderr, "SQL: ", log.Flags())
// 		reformdb = reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
// 	}
// }

// func ReformInsert(b *B) {
// 	var m *Model
// 	wrapExecute(b, func() {
// 		initDB()
// 		m = NewModel()
// 	})

// 	for i := 0; i < b.N; i++ {
// 		d := reformdb.Save(m)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	}
// }

// func ReformInsertMulti(b *B) {
// 	var ms []*Model
// 	wrapExecute(b, func() {
// 		initDB()
// 	})

// 	for i := 0; i < b.N; i++ {
// 		ms = make([]*Model, 0, 100)
// 		for i := 0; i < 100; i++ {
// 			ms = append(ms, NewModel())
// 		}
// 		d := reformdb.Create(&ms)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	}
// }

// func ReformUpdate(b *B) {
// 	var m *Model
// 	wrapExecute(b, func() {
// 		initDB()
// 		m = NewModel()
// 		d := reformdb.Create(&m)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	})

// 	for i := 0; i < b.N; i++ {
// 		d := reformdb.Model(&m).Update("age", 20)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	}
// }

// func ReformRead(b *B) {
// 	var m *Model
// 	wrapExecute(b, func() {
// 		initDB()
// 		m = NewModel()
// 		d := reformdb.Create(&m)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	})
// 	for i := 0; i < b.N; i++ {
// 		d := reformdb.First(&m, m.ID)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	}
// }

// func ReformReadSlice(b *B) {
// 	var m *Model
// 	wrapExecute(b, func() {
// 		initDB()
// 		m = NewModel()
// 		for i := 0; i < 100; i++ {
// 			m.ID = 0
// 			d := reformdb.Create(&m)
// 			if d.Error != nil {
// 				fmt.Println(d.Error)
// 				b.FailNow()
// 			}
// 		}
// 	})

// 	for i := 0; i < b.N; i++ {
// 		var models []*Model
// 		d := reformdb.Limit(100).Find(&models)
// 		if d.Error != nil {
// 			fmt.Println(d.Error)
// 			b.FailNow()
// 		}
// 	}
// }
