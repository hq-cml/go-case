package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//插入
func Insert(addresspwd, dbname, stmt string, val []interface{}) error{
	//db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	//dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8", addresspwd, dbname)
	dsn := fmt.Sprintf("%s/%s?charset=utf8", addresspwd, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		//panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
		return err
	}
	defer db.Close()

	// Prepare statement for inserting data
	//stmtIns, err := db.Prepare("INSERT INTO test VALUES( ?, ? )") // ? = placeholder
	stmtIns, err := db.Prepare(stmt) // ? = placeholder
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return err
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec(val...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return err
	}
	return nil
}

//读取
func Select(addresspwd, dbname, stmt string) error {
	//db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	//dsn := fmt.Sprintf("root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8", dbname)
	dsn := fmt.Sprintf("%s/%s?charset=utf8", addresspwd, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		//panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
		return err
	}
	defer db.Close()

	//rows, err := db.Query("select id, username from user where id = ?", 1)
	rows, err := db.Query(stmt)
	if err != nil {
		return err
	}

	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		fmt.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
