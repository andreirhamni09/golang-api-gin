package handlers

import (
	"api-gin/connections"
	"api-gin/structs"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {

	var dbUsers structs.Users

	if err := c.ShouldBindJSON(&dbUsers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"data":    nil,
			"message": "Lengkapi Data",
		})
		return
	}

	body := c.Request.Body
	payloads, _ := ioutil.ReadAll(body)

	json.Unmarshal(payloads, &dbUsers)
	res := structs.Results{Code: 500, Data: dbUsers, Message: "Unknown Error"}

	switch dbUsers.Role {
	case "0":
		dbUsers.Role = "user"
	case "1":
		dbUsers.Role = "admin"
	default:
		dbUsers.Role = "invalid"
		res.Code = 400
		res.Message = "Invalid User Role"
	}

	if dbUsers.Role != "invalid" {
		genPass, err := EncriptPassword(dbUsers.Password)
		ReturnCheckError(c, err)
		dbUsers.Password = genPass

		if err := connections.DB.Create(&dbUsers).Error; err != nil {
			res.Data = dbUsers
			res.Code = 400
			res.Message = err.Error()
		} else {
			res.Data = dbUsers
			res.Code = 200
			res.Message = "Add new user successfully"
		}
	}
	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}
	ReturnResult(c, res.Code, ginDetail)
}

func GetUsersLimit(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit < 1 {
		limit = 10
	}

	offset, _ := strconv.Atoi(c.Query("offset"))

	if offset < 1 {
		offset = 0
	}

	dbUsers := []structs.Users{}

	if err := connections.DB.Limit(limit).Offset(offset).Find(&dbUsers).Error; err != nil {
		ReturnCheckError(c, err)
	}

	res := structs.Results{Code: 200, Data: dbUsers, Message: "User has successfully retrieve"}
	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	ReturnResult(c, res.Code, ginDetail)
}

func GetUserId(c *gin.Context) {
	id := c.Param("id")

	dbUsers := structs.Users{}

	if err := connections.DB.First(&dbUsers, id).Error; err != nil {
		ReturnCheckError(c, err)
	}
	res := structs.Results{Code: 200, Data: dbUsers, Message: "Users Ditemukan"}

	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	ReturnResult(c, res.Code, ginDetail)
}

func UpdateUserById(c *gin.Context) {
	body := c.Request.Body
	payloads, _ := ioutil.ReadAll(body)

	id := c.Param("id")

	var dbUsers structs.Users

	res := structs.Results{Code: 500, Data: dbUsers, Message: "Unknown Error"}

	if err := connections.DB.First(&dbUsers, id).Error; err != nil {
		ReturnCheckError(c, err)
	}
	json.Unmarshal(payloads, &dbUsers)

	switch dbUsers.Role {
	case "0":
		dbUsers.Role = "user"
	case "1":
		dbUsers.Role = "admin"
	default:
		dbUsers.Role = "invalid"
		res.Code = 400
		res.Message = "Invalid User Role"
	}

	if dbUsers.Role != "invalid" {
		genPass, err := EncriptPassword(dbUsers.Password)
		ReturnCheckError(c, err)
		dbUsers.Password = genPass
		if err := connections.DB.Model(&dbUsers).Updates(&dbUsers).Error; err != nil {
			ReturnCheckError(c, err)
		}
		if !dbUsers.Status {
			connections.DB.Model(&dbUsers).Updates(map[string]interface{}{"status": false})
		}
		res.Data = dbUsers
		res.Code = 200
		res.Message = "Update user successfully"
		ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}
		ReturnResult(c, res.Code, ginDetail)
	}

}

func DeleteUserById(c *gin.Context) {
	id := c.Param("id")

	dbUsers := structs.Users{}

	if err := connections.DB.First(&dbUsers, id).Error; err != nil {
		ReturnCheckError(c, err)
	}
	res := structs.Results{Code: 200, Data: dbUsers, Message: "Users Ditemukan"}

	if err := connections.DB.Delete(&dbUsers).Error; err != nil {
		ReturnCheckError(c, err)
	}

	res.Code = 200
	res.Data = dbUsers
	res.Message = "Berhasil Menghapus Users"

	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}

	ReturnResult(c, res.Code, ginDetail)

}

func LoginUser(c *gin.Context) {
	var dbUser structs.Users
	var userLogin structs.UsersLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body := c.Request.Body
	payloads, _ := ioutil.ReadAll(body)
	json.Unmarshal(payloads, &userLogin)

	res := structs.Results{Code: 200, Data: dbUser, Message: "Gagal Login"}
	connections.DB.Where("username = ?", &userLogin.Username).Find(&dbUser)

	if CekPassword(userLogin.Password, dbUser.Password) {
		res.Data = dbUser
		res.Code = 200
		res.Message = "Berhasil Login"
	}
	ginDetail := gin.H{"code": res.Code, "data": res.Data, "message": res.Message}
	ReturnResult(c, res.Code, ginDetail)
}

/* func LoginUser(c *gin.Context) {
	// body := c.Request.Body
	// payloads, _ := ioutil.ReadAll(body)

	var userLogin structs.UsersLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userLogin.Username != "manu" || userLogin.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
} */
