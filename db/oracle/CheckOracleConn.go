package oracle

import (
	"database/sql"
	"os"
)

func CheckOracleConn(dsn string) error {
	os.Setenv("NLS_LANG", "")
	//if len(os.Args) != 2 {
	//	log.Fatalln(os.Args[0] + " user/password@host:port/sid")
	//}

	db, err := sql.Open("oci8", dsn)
	//fmt.Printf("%+v\n",db)
	if err != nil {
		return err
	}

	defer db.Close()

	rows, err := db.Query("select 3.14 from dual")
	if err != nil {
		return err
	}
	//fmt.Println(rows.Next())
	defer rows.Close()
	return nil
}
