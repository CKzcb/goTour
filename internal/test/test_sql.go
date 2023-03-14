/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: test_sql
 * @Version: 1.0.0
 * @Date: 2023/3/14 12:33
 */

package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/information_schema")
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := db.Query("SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT FROM information_schema.COLUMNS WHERE TABLE_SCHEMA= 'ckTour' AND TABLE_NAME = 'tb_user';")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("all rows: ", rows)
	defer rows.Close()
	for rows.Next() {
		log.Println(rows.Next())
	}

}
