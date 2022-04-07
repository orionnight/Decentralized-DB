package database

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GetJson(c *gin.Context) map[string]interface{} {
	// Create Json object from POST body
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(1, err)
	}
	var jsonObj map[string]interface{}
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		c.AbortWithError(1, err)
	}
	return jsonObj
}
