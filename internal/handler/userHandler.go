package handler

import (
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/models"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/users"
	data "github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"reflect"
)

func ConfigureHandleGetUsers(db data.Db) func(params users.GetUsersParams) middleware.Responder {
	return func(params users.GetUsersParams) middleware.Responder {
		var userList []models.User

		err := db.FindAll(&userList, []string{"Posts"})

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleGetUsers")
			return &users.GetUsersInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		var userPtrs []*models.User
		for i := 0; i < len(userList); i++ {
			userPtrs = append(userPtrs, &userList[i])
		}
		return &users.GetUsersOK{Payload: userPtrs}
	}
}

func ConfigureHandleGetUserById(db data.Db) func(params users.GetUsersIDParams) middleware.Responder {
	return func(params users.GetUsersIDParams) middleware.Responder {
		var user models.User
		err := db.FindById(&user, uint(params.ID), []string{"Posts"})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"user_id": params.ID}).Info("Not found in ConfigureHandleGetUserById")
			return &users.GetUsersIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "User not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleGetUserById")
			return &users.GetUsersIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &users.GetUsersIDOK{Payload: &user}
	}
}

func ConfigureHandleDeleteUser(db data.Db) func(params users.DeleteUsersIDParams) middleware.Responder {
	return func(params users.DeleteUsersIDParams) middleware.Responder {
		var user models.User
		err := db.FindById(&user, uint(params.ID), []string{"Posts"})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"user_id": params.ID}).Info("Not found in ConfigureHandleDeleteUser")
			return &users.DeleteUsersIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "User not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleDeleteUser")
			return &users.DeleteUsersIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		err = db.Delete(user)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleDeleteUser")
			return &users.DeleteUsersIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &users.DeleteUsersIDOK{Payload: &user}
	}
}

func ConfigureHandleCreateUser(db data.Db) func(params users.PostUsersParams) middleware.Responder {
	return func(params users.PostUsersParams) middleware.Responder {
		user := models.User{
			Age:     params.User.Age,
			Email:   params.User.Email,
			IsAdmin: params.User.IsAdmin,
			Name:    params.User.Name,
		}
		err := db.Create(&user)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleCreateUser")
			return &users.PostUsersInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &users.PostUsersOK{Payload: &(user)}
	}
}

func ConfigureHandleUpdateUser(db data.Db) func(params users.PutUsersIDParams) middleware.Responder {
	return func(params users.PutUsersIDParams) middleware.Responder {
		var user models.User
		err := db.FindById(&user, uint(params.ID), []string{"Posts"})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"user_id": params.ID}).Info("Not found in ConfigureHandleUpdateUser")
			return &users.PutUsersIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "User not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleUpdateUser")
			return &users.PutUsersIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		if params.User.Age != nil {
			user.Age = params.User.Age
		}

		if params.User.Email != "" {
			user.Email = params.User.Email
		}

		// Needed to determine if boolean has been set
		if reflect.ValueOf(params.User.IsAdmin).Bool() {
			user.IsAdmin = params.User.IsAdmin
		}

		if params.User.Name != "" {
			user.Name = params.User.Name
		}

		err = db.Update(&user)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleCreateUser")
			return &users.PutUsersIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &users.PutUsersIDOK{Payload: &user}
	}
}
