package server

import(
	"database/sql"
	_"mysql_master"
)

func Koneksi()(*sql.DB,error){
	db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/joopizza")

	if err != nil{
		return nil , err
	}

	return db , nil
}