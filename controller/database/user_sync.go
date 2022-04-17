package database

import (
	"encoding/json"
	"fmt"
)

func SyncDBCreation(user string) {
	println("Start sync db create...")

	data := User{}

	err := json.Unmarshal([]byte(user), &data)
	if err != nil {
		fmt.Println(err.Error())
		//json: Unmarshal(non-pointer main.Request)
	}
	_, restErr := CreateUser(&data)
	if restErr != nil {
		fmt.Printf(fmt.Sprint(restErr.Status), restErr)
		return
	}
	println("End sync db create, success!\n")
}

// func SyncDBCreation( map[string]interface{}) {
// 	println("Start sync db create...")

// 	data := User{}

// 	err := json.Unmarshal([]byte(user), &data)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		//json: Unmarshal(non-pointer main.Request)
// 	}
// 	_, restErr := CreateUser(&data)
// 	if restErr != nil {
// 		fmt.Printf(fmt.Sprint(restErr.Status), restErr)
// 		return
// 	}
// 	println("End sync db create, success!\n")
// }

func SyncDBDeletion(email string) {
	println("Start sync db deletion...")

	if email == "" {
		restErr := BadRequest("no email.")
		print(restErr)
		return
	}
	restErr := DeleteUser(email)
	if restErr != nil {
		print(restErr)
		return
	}
	println("End sync db deletion, success!\n")
}

func SyncDBUpdate(dbUpdate string) {
	println("Start sync db update...")

	data := UpdateDB{}
	json.Unmarshal([]byte(dbUpdate), &data)

	err := json.Unmarshal([]byte(dbUpdate), data)
	if err != nil {
		fmt.Println(err.Error())
	}

	userEmail := data.Email
	field := data.Field
	value := data.Value
	if userEmail == "" {
		restErr := BadRequest("no email.")
		println(restErr)
		return
	}
	if field == "" {
		restErr := BadRequest("no field.")
		println(restErr)
		return
	}
	if value == "" {
		restErr := BadRequest("no value.")
		println(restErr)
		return
	}
	_, restErr := UpdateUser(userEmail, field, value)
	if restErr != nil {
		println(restErr)
		return
	}
	println("End sync db update, success!\n")
}
