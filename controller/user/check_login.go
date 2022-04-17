package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUserLogin(c *gin.Context) bool {
	if bcuser == nil {
		fmt.Println("User is not logged in. Please login first!")
		c.String(http.StatusForbidden, "User is not logged in. Please login first!")
		return false
	}
	return true
}
