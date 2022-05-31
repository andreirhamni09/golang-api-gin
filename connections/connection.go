package connections

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"		

	"os"
	"github.com/joho/godotenv"
)

var (
	DB  *gorm.DB
	Err error
)

func Connection() {
	load_env := godotenv.Load(".env")
	if load_env != nil {
		fmt.Println("Gagal Load Environment", load_env)
	}
	
	db_conn := os.Getenv("DB_CONNECTION")
	db_host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_DATABASE")
	db_user := os.Getenv("DB_USERNAME")
	db_pass := os.Getenv("DB_PASSWORD")

	str_connection := db_user + ":"+ db_pass + "@(" + db_host + ")/" + db_name + "?charset=utf8&parseTime=True&loc=Local"

	DB, Err = gorm.Open(db_conn, str_connection)
	if Err != nil {
		fmt.Println("Gagal Koneksi", Err)
	} else {
		fmt.Println("Berhasil Koneksi")
	}
	/* DB.AutoMigrate(structs.Users{})
	DB.AutoMigrate(structs.Products{}) */
}
