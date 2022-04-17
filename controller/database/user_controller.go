package database

import (
	"encoding/json"
	"fmt"
	"net/http"

	blockchain_user "example.com/ece1770/controller/user"
	"example.com/ece1770/transaction"
	"github.com/gin-gonic/gin"
)

// http://142.150.199.223/db/ping
func Ping(c *gin.Context) {
	fmt.Println("Start db ping...")
	c.String(http.StatusOK, "pong")
}

func HandleDBCreation(c *gin.Context) {
	if !blockchain_user.CheckUserLogin(c) {
		return
	}

	fmt.Println("Start db create...")
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)

	jsonResp, err := json.Marshal(newUser)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	res := map[string]interface{}{
		"op":  "CREATE",
		"arg": string(jsonResp),
	}
	target, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error appending op to JSON. Err: %s", err)
		return
	}

	transaction.SendRawTransaction(string(target), blockchain_user.GetPrivateKeyHex())

	// test sync
	// SyncDBCreation(string(target))

	println("End db create, success!\n")
}

// http://142.150.199.223/db/delete?email=abby123@gmail.com
func HandleDBDeletion(c *gin.Context) {
	if !blockchain_user.CheckUserLogin(c) {
		return
	}
	fmt.Println("Start db deletion...")
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	restErr := DeleteUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})

	res := map[string]interface{}{
		"op":  "DELETE",
		"arg": string(userEmail),
	}
	target, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error appending op to JSON. Err: %s", err)
		return
	}

	transaction.SendRawTransaction(string(target), blockchain_user.GetPrivateKeyHex())

	// test sync
	// SyncDBDeletion(string(target))

	println("End db delete, success!\n")
}

// http://142.150.199.223/db/update?email=abby123@gmail.com&field=name&value=Rachel
func HandleDBUpdate(c *gin.Context) {
	if !blockchain_user.CheckUserLogin(c) {
		return
	}
	fmt.Println("Start db update...")
	userEmail := c.Query("email")
	field := c.Query("field")
	value := c.Query("value")
	if userEmail == "" {
		restErr := BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if field == "" {
		restErr := BadRequest("no field.")
		c.JSON(restErr.Status, restErr)
		return
	}
	if value == "" {
		restErr := BadRequest("no value.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := UpdateUser(userEmail, field, value)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)

	updateDB := UpdateDB{Email: userEmail, Field: field, Value: value}
	jsonResp, err := json.MarshalIndent(updateDB, "", "\t")
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}

	res := map[string]interface{}{
		"op":  "UPDATE",
		"arg": string(jsonResp),
	}
	target, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error appending op to JSON. Err: %s", err)
		return
	}

	transaction.SendRawTransaction(string(target), blockchain_user.GetPrivateKeyHex())

	// test sync
	// SyncDBUpdate(string(target))

	println("End db update, success!\n")
}

// http://142.150.199.223/db/find?email=abby123@gmail.com
func HandleDBFind(c *gin.Context) {
	if !blockchain_user.CheckUserLogin(c) {
		return
	}
	fmt.Println("Start db find...")
	userEmail := c.Query("email")
	if userEmail == "" {
		restErr := BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := FindUser(userEmail)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
	println("End db find, success!\n")
}
