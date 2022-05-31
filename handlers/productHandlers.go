package handlers

import (
	"api-gin/connections"
	"api-gin/structs"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"api-gin/functions"
)

func CreateProduct(c *gin.Context) {
	body := c.Request.Body
	payloads, _ := ioutil.ReadAll(body)

	var dbProducts structs.Products
	json.Unmarshal(payloads, &dbProducts)

	res := structs.Results{Code: 500, Data: dbProducts, Message: "Unknown Error"}

	if err := connections.DB.Create(&dbProducts).Error; err != nil {
		ginDetail := gin.H{"code": 400, "data": dbProducts, "message": err.Error()}
		functions.ReturnResult(c, 400, ginDetail)
	}
	res.Data = dbProducts
	res.Code = 200
	res.Message = "Add new Product successfully"

	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}
	functions.ReturnResult(c, res.Code, ginDetail)
}
func GetProductsLimit(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 1 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.Query("offset"))

	if offset < 1 {
		offset = 0
	}

	dbProducts := []structs.Products{}

	if err := connections.DB.Limit(limit).Offset(offset).Find(&dbProducts).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}

	res := structs.Results{Code: 200, Data: dbProducts, Message: "Products has successfully retrieve"}
	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	functions.ReturnResult(c, res.Code, ginDetail)
}
func GetProductId(c *gin.Context) {
	id := c.Param("id")

	dbProducts := structs.Products{}

	if err := connections.DB.First(&dbProducts, id).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}
	res := structs.Results{Code: 200, Data: dbProducts, Message: "Produk Ditemukan"}

	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	functions.ReturnResult(c, res.Code, ginDetail)
}
func UpdateProductById(c *gin.Context) {
	body := c.Request.Body
	payloads, _ := ioutil.ReadAll(body)

	id := c.Param("id")

	var dbProducts structs.Products

	res := structs.Results{Code: 500, Data: dbProducts, Message: "Unknown Error"}

	if err := connections.DB.First(&dbProducts, id).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}
	json.Unmarshal(payloads, &dbProducts)

	if err := connections.DB.Model(&dbProducts).Updates(&dbProducts).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}
	if !dbProducts.Status {
		connections.DB.Model(&dbProducts).Updates(map[string]interface{}{"status": false})
	}
	res.Data = dbProducts
	res.Code = 200
	res.Message = "Update Product successfully"
	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}
	functions.ReturnResult(c, res.Code, ginDetail)

}
func DeletProductById(c *gin.Context) {
	id := c.Param("id")

	dbProducts := structs.Products{}

	res := structs.Results{Code: 500, Data: dbProducts, Message: "Unknown Error"}

	if err := connections.DB.First(&dbProducts, id).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}
	res.Code = 200
	res.Data = dbProducts
	res.Message = "Product Ditemukan"

	if err := connections.DB.Delete(&dbProducts).Error; err != nil {
		functions.ReturnCheckError(c, err)
	}

	res.Code = 200
	res.Data = dbProducts
	res.Message = "Berhasil Menghapus Product"

	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	functions.ReturnResult(c, res.Code, ginDetail)
}
