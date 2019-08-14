package handler

import (
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/models"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/posts"
	data "github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
)

func ConfigureHandleGetPosts(db data.Db) func(params posts.GetPostsParams) middleware.Responder {
	return func(params posts.GetPostsParams) middleware.Responder {
		var post []models.Post

		err := db.FindAll(&post, []string{})

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleGetPosts")
			return &posts.GetPostsInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		var postPtrs []*models.Post
		for i := 0; i < len(post); i++ {
			postPtrs = append(postPtrs, &post[i])
		}
		return &posts.GetPostsOK{Payload: postPtrs}
	}
}

func ConfigureHandleGetPostById(db data.Db) func(params posts.GetPostsIDParams) middleware.Responder {
	return func(params posts.GetPostsIDParams) middleware.Responder {
		var post models.Post
		err := db.FindById(&post, uint(params.ID), []string{})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"post_id": params.ID}).Info("Not found in ConfigureHandleGetPostById")
			return &posts.GetPostsIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "Post not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleGetPostById")
			return &posts.GetPostsIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &posts.GetPostsIDOK{Payload: &post}
	}
}

func ConfigureHandleDeletePost(db data.Db) func(params posts.DeletePostsIDParams) middleware.Responder {
	return func(params posts.DeletePostsIDParams) middleware.Responder {
		var post models.Post
		err := db.FindById(&post, uint(params.ID), []string{})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"post_id": params.ID}).Info("Not found in ConfigureHandleDeletePost")
			return &posts.DeletePostsIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "Post not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleDeletePost")
			return &posts.DeletePostsIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		err = db.Delete(post)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleDeletePost")
			return &posts.DeletePostsIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &posts.DeletePostsIDOK{Payload: &post}
	}
}

func ConfigureHandleCreatePost(db data.Db) func(params posts.PostPostsParams) middleware.Responder {
	return func(params posts.PostPostsParams) middleware.Responder {
		post := models.Post{
			Description: params.Post.Description,
			UserID:      params.Post.UserID,
		}
		err := db.Create(&post)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleCreatePost")
			return &posts.PostPostsInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &posts.PostPostsOK{Payload: &(post)}
	}
}

func ConfigureHandleUpdatePost(db data.Db) func(params posts.PutPostsIDParams) middleware.Responder {
	return func(params posts.PutPostsIDParams) middleware.Responder {
		var post models.Post
		err := db.FindById(&post, uint(params.ID), []string{})

		if err != nil && err.Error() == "record not found" {
			log.WithFields(log.Fields{"post_id": params.ID}).Info("Not found in ConfigureHandleUpdatePost")
			return &posts.PutPostsIDNotFound{Payload: &models.NotFound{
				Code:    404,
				Message: "Post not found",
			}}
		}

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleUpdatePost")
			return &posts.PutPostsIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		if params.Post.UserID != nil {
			post.UserID = params.Post.UserID
		}

		if params.Post.Description != "" {
			post.Description = params.Post.Description
		}

		err = db.Update(&post)

		if err != nil {
			log.WithError(err).Error("Error in ConfigureHandleUpdatePost")
			return &posts.PutPostsIDInternalServerError{Payload: &models.Error{
				Code:    500,
				Message: err.Error(),
			}}
		}

		return &posts.PutPostsIDOK{Payload: &post}
	}
}
