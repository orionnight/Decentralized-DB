package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user *User) (*User, *RestErr) {
	user, restErr := Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func FindUser(email string) (*User, *RestErr) {
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func DeleteUser(email string) *RestErr {
	restErr := Delete(email)
	if restErr != nil {
		return restErr
	}
	return nil
}

func UpdateUser(email string, field string, value string) (*User, *RestErr) {
	user, restErr := Update(email, field, value)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func Create(user *User) (*User, *RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		restErr := InternalErr("can't insert user to the ")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func Find(email string) (*User, *RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

func Delete(email string) *RestErr {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.DeleteOne(ctx, bson.M{"email": email})
	if err != nil {
		restErr := NotFound("faild to delete.")
		return restErr
	}
	if result.DeletedCount == 0 {
		restErr := NotFound("user not found.")
		return restErr
	}
	return nil
}

func Update(email string, field string, value string) (*User, *RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
