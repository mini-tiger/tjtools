package oracle

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"sync"
)

type OrclData struct {
	sync.Mutex
	DB *sql.DB
}

func NewOrclClient(dsn string) (err error,db *sql.DB) {
	//db, err := sql.Open("oci8", "test1/test1@1.119.132.155:1521/orcl")
	db, err = sql.Open("oci8", dsn)
	//fmt.Printf("%+v\n",db)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return

}



func (o *OrclData)GetSqlData(sql string)  (err error,rows *sql.Rows){
	o.Lock()
	defer func() {
		o.Unlock()
	}()

	rows, err = o.DB.Query(sql)
	if err != nil {
		return err,rows
	}
	return

}

//func main() {
//	os.Setenv("NLS_LANG", "")
//	//if len(os.Args) != 2 {
//	//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
//	//}
//
//	db, err := sql.Open("oci8", "test1/test1@1.119.132.155:1521/orcl")
//	//fmt.Printf("%+v\n",db)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	defer db.Close()
//
//	rows, err := db.Query("select train_serial from tf_op_train where rownum < 2")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	defer rows.Close()
//
//	for rows.Next() {
//		var data string
//		rows.Scan(&data)
//		fmt.Println(data)
//	}
//	if err = rows.Err(); err != nil {
//		log.Fatalln(err)
//	}
//}
