package handler

import (
	"errors"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/posts"
	data "github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
	"reflect"
	"testing"
)

func TestConfigureHandleGetPostById(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleGetPostById(db)
		params := posts.GetPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.GetPostsIDNotFound{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected not found type but got something else")
		}

		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
	})

	t.Run("Error", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleGetPostById(db)
		params := posts.GetPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.GetPostsIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
	})

	t.Run("Success", func(t *testing.T) {
		db := data.CreateTestDb()
		modelId := uint(5)
		db.FindByIdHandler = func(model interface{}, id uint, loads []string) error {
			if id != modelId {
				t.Errorf("Wrong id passed in handler")
			}
			if len(loads) != 0 {
				t.Errorf("Should not have loaded anything")
			}

			return nil
		}

		handler := ConfigureHandleGetPostById(db)
		params := posts.GetPostsIDParams{
			ID: float64(modelId),
		}
		expectedType := reflect.TypeOf(&posts.GetPostsIDOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}

		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
	})
}

func TestConfigureHandleGetPosts(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindAllHandler = func(interface{}, []string) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleGetPosts(db)
		params := posts.GetPostsParams{}
		expectedType := reflect.TypeOf(&posts.GetPostsInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.FindAllCount != 1 {
			t.Errorf("Should have called get once")
		}
	})

	t.Run("Success", func(t *testing.T) {
		db := data.CreateTestDb()

		handler := ConfigureHandleGetPosts(db)
		params := posts.GetPostsParams{}
		expectedType := reflect.TypeOf(&posts.GetPostsOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}

		if db.FindAllCount != 1 {
			t.Errorf("Should have called get once")
		}
	})
}

func TestConfigureHandleCreatePost(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		db := data.CreateTestDb()
		db.CreateHandler = func(interface{}) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleCreatePost(db)
		params := posts.PostPostsParams{}
		expectedType := reflect.TypeOf(&posts.PostPostsInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.CreateCount != 1 {
			t.Errorf("Should have called create once")
		}
	})

	t.Run("Success", func(t *testing.T) {
		db := data.CreateTestDb()
		handler := ConfigureHandleCreatePost(db)
		params := posts.PostPostsParams{}
		expectedType := reflect.TypeOf(&posts.PostPostsOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.CreateCount != 1 {
			t.Errorf("Should have called create once")
		}
	})
}

func TestConfigureHandleDeletePost(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleDeletePost(db)
		params := posts.DeletePostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.DeletePostsIDNotFound{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Error(reflect.TypeOf(res))
			t.Errorf("Expected not found type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.DeleteCount != 0 {
			t.Errorf("Should not have called delete")
		}
	})

	t.Run("Error In Find", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleDeletePost(db)
		params := posts.DeletePostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.DeletePostsIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.DeleteCount != 0 {
			t.Errorf("Should not have called delete")
		}
	})

	t.Run("Error In Delete", func(t *testing.T) {
		db := data.CreateTestDb()
		db.DeleteHandler = func(interface{}) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleDeletePost(db)
		params := posts.DeletePostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.DeletePostsIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.DeleteCount != 1 {
			t.Errorf("Should have called delete once")
		}
	})

	t.Run("Success", func(t *testing.T) {
		db := data.CreateTestDb()
		handler := ConfigureHandleDeletePost(db)
		params := posts.DeletePostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.DeletePostsIDOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.DeleteCount != 1 {
			t.Errorf("Should have called delete once")
		}
	})
}

func TestConfigureHandleUpdatePost(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleUpdatePost(db)
		params := posts.PutPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.PutPostsIDNotFound{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Error(reflect.TypeOf(res))
			t.Errorf("Expected not found type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.UpdateCount != 0 {
			t.Errorf("Should not have called update")
		}
	})

	t.Run("Error In Find", func(t *testing.T) {
		db := data.CreateTestDb()
		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleUpdatePost(db)
		params := posts.PutPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.PutPostsIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.UpdateCount != 0 {
			t.Errorf("Should not have called update")
		}
	})

	t.Run("Error In Update", func(t *testing.T) {
		db := data.CreateTestDb()
		db.UpdateHandler = func(interface{}) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleUpdatePost(db)
		params := posts.PutPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.PutPostsIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.UpdateCount != 1 {
			t.Errorf("Should have called update once")
		}
	})

	t.Run("Success", func(t *testing.T) {
		db := data.CreateTestDb()
		handler := ConfigureHandleUpdatePost(db)
		params := posts.PutPostsIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&posts.PutPostsIDOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}
		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
		if db.UpdateCount != 1 {
			t.Errorf("Should have called update once")
		}
	})
}
