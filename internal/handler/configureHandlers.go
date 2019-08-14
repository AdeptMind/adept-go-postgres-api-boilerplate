package handler

import (
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/posts"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/users"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
)

func ConfigureApiHandlers(api *operations.BoilerplateAPI, db db.Db) {
	configurePostApiHandlers(api, db)
	configureUserApiHandlers(api, db)
}

func configurePostApiHandlers(api *operations.BoilerplateAPI, db db.Db) {
	api.PostsGetPostsHandler = posts.GetPostsHandlerFunc(ConfigureHandleGetPosts(db))
	api.PostsGetPostsIDHandler = posts.GetPostsIDHandlerFunc(ConfigureHandleGetPostById(db))
	api.PostsDeletePostsIDHandler = posts.DeletePostsIDHandlerFunc(ConfigureHandleDeletePost(db))
	api.PostsPostPostsHandler = posts.PostPostsHandlerFunc(ConfigureHandleCreatePost(db))
	api.PostsPutPostsIDHandler = posts.PutPostsIDHandlerFunc(ConfigureHandleUpdatePost(db))
}

func configureUserApiHandlers(api *operations.BoilerplateAPI, db db.Db) {
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(ConfigureHandleGetUsers(db))
	api.UsersGetUsersIDHandler = users.GetUsersIDHandlerFunc(ConfigureHandleGetUserById(db))
	api.UsersDeleteUsersIDHandler = users.DeleteUsersIDHandlerFunc(ConfigureHandleDeleteUser(db))
	api.UsersPostUsersHandler = users.PostUsersHandlerFunc(ConfigureHandleCreateUser(db))
	api.UsersPutUsersIDHandler = users.PutUsersIDHandlerFunc(ConfigureHandleUpdateUser(db))
}
