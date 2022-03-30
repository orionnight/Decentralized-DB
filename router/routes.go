package router

import (
	"fmt"
	"net/http"

	"example.com/ece1770/services"
)

func DBCreate(w http.ResponseWriter, r *http.Request, json_obj map[string]interface{}) {
	fmt.Fprintf(w, "Json object is %v\n", json_obj)
}

func CreateKeystore(w http.ResponseWriter, r *http.Request, json_obj map[string]interface{}) {
	if password, ok := json_obj["password"].(string); ok {
		services.CreateKS(password)
	} else {
		fmt.Fprintf(w, "password not provided!")
	}

}
