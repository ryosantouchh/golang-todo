package todo

import (
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

// GormStore implement storer ==> can be use with
func (s *GormStore) New(todo *Todo) error {
	return s.db.Create(todo).Error
}
