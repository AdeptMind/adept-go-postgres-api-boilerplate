package db

type Test struct {
	FindAllCount    int
	FindAllHandler  func(interface{}, []string) error
	FindByIdCount   int
	FindByIdHandler func(interface{}, uint, []string) error
	DeleteCount     int
	DeleteHandler   func(interface{}) error
	CreateCount     int
	CreateHandler   func(interface{}) error
	UpdateCount     int
	UpdateHandler   func(interface{}) error
}

func (t *Test) FindAll(model interface{}, loads []string) error {
	t.FindAllCount++
	return t.FindAllHandler(model, loads)
}

func (t *Test) FindById(model interface{}, id uint, loads []string) error {
	t.FindByIdCount++
	return t.FindByIdHandler(model, id, loads)
}

func (t *Test) Delete(model interface{}) error {
	t.DeleteCount++
	return t.DeleteHandler(model)
}

func (t *Test) Create(model interface{}) error {
	t.CreateCount++
	return t.CreateHandler(model)
}

func (t *Test) Update(model interface{}) error {
	t.UpdateCount++
	return t.UpdateHandler(model)
}

func CreateTestDb() *Test {
	return &Test{
		FindAllCount:    0,
		FindAllHandler:  func(interface{}, []string) error { return nil },
		FindByIdCount:   0,
		FindByIdHandler: func(interface{}, uint, []string) error { return nil },
		DeleteCount:     0,
		DeleteHandler:   func(interface{}) error { return nil },
		CreateCount:     0,
		CreateHandler:   func(interface{}) error { return nil },
		UpdateCount:     0,
		UpdateHandler:   func(interface{}) error { return nil },
	}
}
