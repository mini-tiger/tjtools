package oracle

import (
	"database/sql"
)


func ConnKeepLive(db *sql.DB) (err error) {
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