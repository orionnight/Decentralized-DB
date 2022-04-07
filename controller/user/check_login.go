package user

import "fmt"

func CheckUserLogin() bool {
	if bcuser == nil {
		fmt.Println("User is not logged in. Please login first!")
		return false
	}
	return true
}
