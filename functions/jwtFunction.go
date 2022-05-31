package functions
import (
	"github.com/gin-gonic/gin"
	
	"os"
	"github.com/joho/godotenv"
)
func GetJwtKey() (string, error) {
	load_env := godotenv.Load(".env")
	if load_env != nil {
		panic(load_env)
	}
	keyApi := os.Getenv("APP_API_KEY")
	return keyApi, nil
}

func GenerateJwtToken() (string){
	keyApi, _ := GetJwtKey()
	return keyApi
}

func TampilkanToken(c *gin.Context) {
	GToken := GenerateJwtToken()
	c.JSON(200, gin.H{"Token" : GToken})
}