package oracle

import (
	"database/sql"
	"os"
)

func CheckOracleConn(dsn *string) (err error) {
	var db *sql.DB
	var rows *sql.Rows

	// 每次生成的内存地址不同，不能重复利用
	defer func() {
		db = nil
		rows = nil
	}()
	_ = os.Setenv("NLS_LANG", "")
	//if len(os.Args) != 2 {
	//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
	//}

	db, err = sql.Open("oci8", *dsn)
	//fmt.Printf("%+v\n",db)
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()

	rows, err = db.Query("select 3.14 from dual")
	if err != nil {
		return err
	}
	//fmt.Println(rows.Next())
	defer func() {
		_ = rows.Close()
	}()
	return nil
}

func CheckOracleLive(db *sql.DB) (err error) {
	var rows *sql.Rows
	// 每次生成的内存地址不同，不能重复利用
	defer func() {
		rows = nil
	}()

	rows, err = db.Query("select 3.14 from dual")
	if err != nil {
		return err
	}
	//fmt.Println(rows.Next())
	defer func() {
		_ = rows.Close()
	}()
	return nil
}