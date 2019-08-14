package db

type Db interface {
	FindAll(model interface{}, loads []string) error
	FindById(model interface{}, id uint, loads []string) error
	Delete(model interface{}) error
	Create(model interface{}) error
	Update(model interface{}) error
}
