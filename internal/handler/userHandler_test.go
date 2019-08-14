package handler

import (
	"errors"
	"github.com/adeptmind/adept-go-postgres-api-boilerplate/gen/restapi/operations/users"
	data "github.com/adeptmind/adept-go-postgres-api-boilerplate/internal/db"
	"reflect"
	"testing"
)

func TestConfigureHandleGetUserById(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()

		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleGetUserById(db)
		params := users.GetUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.GetUsersIDNotFound{})
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

		handler := ConfigureHandleGetUserById(db)
		params := users.GetUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.GetUsersIDInternalServerError{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
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
			if len(loads) != 1 || loads[0] != "Posts" {
				t.Errorf("Should have only loaded posts")
			}

			return nil
		}

		handler := ConfigureHandleGetUserById(db)
		params := users.GetUsersIDParams{
			ID: float64(modelId),
		}
		expectedType := reflect.TypeOf(&users.GetUsersIDOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}

		if db.FindByIdCount != 1 {
			t.Errorf("Should have called get by id once")
		}
	})
}

func TestConfigureHandleGetUsers(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		db := data.CreateTestDb()

		db.FindAllHandler = func(interface{}, []string) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleGetUsers(db)
		params := users.GetUsersParams{}
		expectedType := reflect.TypeOf(&users.GetUsersInternalServerError{})
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

		handler := ConfigureHandleGetUsers(db)
		params := users.GetUsersParams{}
		expectedType := reflect.TypeOf(&users.GetUsersOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected success type but got something else")
		}

		if db.FindAllCount != 1 {
			t.Errorf("Should have called get once")
		}
	})
}

func TestConfigureHandleCreateUser(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		db := data.CreateTestDb()

		db.CreateHandler = func(interface{}) error {
			return errors.New("some other error")
		}

		handler := ConfigureHandleCreateUser(db)
		params := users.PostUsersParams{}
		expectedType := reflect.TypeOf(&users.PostUsersInternalServerError{})
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

		handler := ConfigureHandleCreateUser(db)
		params := users.PostUsersParams{}
		expectedType := reflect.TypeOf(&users.PostUsersOK{})
		res := handler(params)

		if reflect.TypeOf(res) != expectedType {
			t.Errorf("Expected 500 error type but got something else")
		}

		if db.CreateCount != 1 {
			t.Errorf("Should have called create once")
		}
	})
}

func TestConfigureHandleDeleteUser(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()

		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleDeleteUser(db)
		params := users.DeleteUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.DeleteUsersIDNotFound{})
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

		handler := ConfigureHandleDeleteUser(db)
		params := users.DeleteUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.DeleteUsersIDInternalServerError{})
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

		handler := ConfigureHandleDeleteUser(db)
		params := users.DeleteUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.DeleteUsersIDInternalServerError{})
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

		handler := ConfigureHandleDeleteUser(db)
		params := users.DeleteUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.DeleteUsersIDOK{})
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

func TestConfigureHandleUpdateUser(t *testing.T) {
	t.Run("Not Found", func(t *testing.T) {
		db := data.CreateTestDb()

		db.FindByIdHandler = func(interface{}, uint, []string) error {
			return errors.New("record not found")
		}

		handler := ConfigureHandleUpdateUser(db)
		params := users.PutUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.PutUsersIDNotFound{})
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

		handler := ConfigureHandleUpdateUser(db)
		params := users.PutUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.PutUsersIDInternalServerError{})
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

		handler := ConfigureHandleUpdateUser(db)
		params := users.PutUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.PutUsersIDInternalServerError{})
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

		handler := ConfigureHandleUpdateUser(db)
		params := users.PutUsersIDParams{
			ID: 0,
		}
		expectedType := reflect.TypeOf(&users.PutUsersIDOK{})
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
