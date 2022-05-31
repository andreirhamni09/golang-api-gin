package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ReturnCheckError(c *gin.Context, err error) {
	if err != nil {
		res := gin.H{"code": http.StatusInternalServerError, "data": nil, "message": err.Error()}
		c.JSON(http.StatusInternalServerError, res)
	}
}

func ReturnResult(c *gin.Context, status int, result interface{}) {
	c.JSON(status, result)
}

func EncriptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	pwd := string(hashedPassword)
	return pwd, err
}
func CekPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
