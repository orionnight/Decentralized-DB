package database

import "go.mongodb.org/mongo-driver/bson/primitive"

// type nosql struct {
// 	Statement string `json:"db_statement"`
// 	PrivateKey string `json:"caller_pk"`
// }

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type UpdateDB struct {
	Email string
	Field string
	Value string
}

// utils
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad request",
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "not found",
	}
}

func InternalErr(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "internal server error",
	}
}
