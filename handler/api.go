package handler

import (
	"echo/server"
	"fmt"
	_ "mysql-master"
	"net/http"

	"github.com/labstack/echo"
)

type menu struct {
	Id_menu    string
	Nama_menu  string
	Deskripsi  string
	Jenis      string
	Harga      string
	Url_gambar string
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()

	return c.JSON(http.StatusOK, data)
}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("insert into tbl_menu values(?,?,?,?,?,?)", nil, nama, deskripsi, url_gambar, jenis, harga)

	if err != nil {
		fmt.Println("Menu gagal ditambahkan")
		return c.JSON(http.StatusOK, "Gagal menambahkan menu")
	} else {
		fmt.Println("Menu berhasil ditambahkan")
		return c.JSON(http.StatusOK, "Berhasil menambahkan menu")
	}

}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("update tbl_menu set nama_menu = ? , deskripsi = ? , harga = ? , jenis = ? , url_gambar = ? where id_menu = ?", nama, deskripsi, harga, jenis, url_gambar, id)

	if err != nil {
		fmt.Println("Menu gagal diubah")
		return c.JSON(http.StatusOK, "Gagal menambahkan menu")
	} else {
		fmt.Println("Menu berhasil ditambahkan")
		return c.JSON(http.StatusOK, "Berhasil mengubah menu")
	}

}

func menu_makanan() {
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
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)

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
