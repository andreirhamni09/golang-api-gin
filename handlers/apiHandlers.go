package handlers

import (
	"api-gin/structs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandlerReq() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", HomePage)

	//Tambah User Baru
	r.POST("/user", CreateUsers)

	//Find User Berdasarkan Limit
	r.GET("/users", GetUsersLimit)

	//Find User Berdasarkan Id
	r.GET("/user/:id", GetUserId)

	//Update User Berdasarkan Id
	r.PUT("/user/:id", UpdateUserById)

	//Hapus User Berdasarkan Id
	r.DELETE("/user/:id", DeleteUserById)

	//Hapus User Berdasarkan Id
	r.POST("/login", LoginUser)

	//Tambah Produt Baru
	r.POST("/product", CreateProduct)

	//Find ProductBerdasarkan Limit
	r.GET("/products", GetProductsLimit)

	//Find Product Berdasarkan I d
	r.GET("/product/:id", GetProductId)

	//Update Product Brdasarkan Id
	r.PUT("/product/:id", UpdateProductById)

	//Hapus Product Berdasarkan Id
	r.DELETE("/product/:id", DeletProductById)

	r.Run(":9000")
}

func HomePage(c *gin.Context) {
	homeRes := structs.Results{}
	homeRes.Code = 200
	homeRes.Data = nil
	homeRes.Message = "Ini Halaman ome"

	ginDetail := gin.H{"code": homeRes.Code, "data": homeRes.Data, "message": homeRes.Message}
	c.JSON(homeRes.Code, ginDetail)
}
