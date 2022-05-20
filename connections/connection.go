package connections

import (
	"api-gin/structs"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB  *gorm.DB
	Err error
)

func Connection() {
	DB, Err = gorm.Open("mysql", "andre:0s7cosUcjNWmOHbF@/andre?charset=utf8&parseTime=True&loc=Local")
	if Err != nil {
		fmt.Println("Gagal Koneksi", Err)
	} else {
		fmt.Println("Berhasil Koneksi")
	}
	DB.AutoMigrate(structs.Users{})
	DB.AutoMigrate(structs.Products{})
}
