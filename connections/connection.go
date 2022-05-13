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
	DB, Err = gorm.Open("mysql", "root:@/db_nasabah_gin?charset=utf8&parseTime=True&loc=Local")
	if Err != nil {
		fmt.Println("Gagal Koneksi", Err)
	} else {
		fmt.Println("Berhasil Koneksi")
	}
	DB.AutoMigrate(structs.Users{})
	DB.AutoMigrate(structs.Products{})
}
