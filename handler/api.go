package handler

import (
	"echo/server"
	"fmt"
	_ "mysql-master"
	"net/http"

	"github.com/labstack/echo"
)

type menu struct {
	IDMenu    string
	NamaMenu  string
	Deskripsi string
	Jenis     string
	Harga     string
	UrlGambar string
}

var data []menu

func BacaData(c echo.Context) error {
	menuMakanan()

	return c.JSON(http.StatusOK, data)
}

func menuMakanan() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tbl_menu")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.IDMenu, &each.NamaMenu, &each.Deskripsi, &each.UrlGambar, &each.Jenis, &each.Harga)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}
