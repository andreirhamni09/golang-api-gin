package handlers

import (
	"api-gin/structs"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func HandlerReq() {
	r := gin.Default()
	r.Use(CORS())
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
