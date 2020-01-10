package oracle

import (
	"database/sql"
	"os"
)
var db *sql.DB
var rows *sql.Rows
func CheckOracleConn(dsn *string) (err error) {
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