package db

import "github.com/jinzhu/gorm"

type Postgres struct {
	db *gorm.DB
}

func (p Postgres) Delete(model interface{}) error {
	return p.db.Delete(model).Error
}

func (p Postgres) Create(model interface{}) error {
	return p.db.Create(model).Error
}

func (p Postgres) Update(model interface{}) error {
	return p.db.Save(model).Error
}

func (p Postgres) FindById(model interface{}, id uint, loads []string) error {
	query := p.db
	for i := 0; i < len(loads); i++ {
		query = query.Preload(loads[i])
	}
	return query.First(model, id).Error
}

func (p Postgres) FindAll(model interface{}, loads []string) error {
	query := p.db
	for i := 0; i < len(loads); i++ {
		query = query.Preload(loads[i])
	}
	return query.Find(model).Error
}

func CreatePostgres(db *gorm.DB) *Postgres {
	return &Postgres{db: db}
}
