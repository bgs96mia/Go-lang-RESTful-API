package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entities *T) error {
	return db.Create(entities).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entities *T) error {
	return db.Save(entities).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entities *T) error {
	return db.Delete(entities).Error
}

func (r *Repository[T]) CountById(db *gorm.DB, username any) (int64, error) {
	var total int64
	err := db.Model(new(T)).Where("username = ?", username).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindById(db *gorm.DB, entities *T, username any) error {
	return db.Where("username = ?", username).Take(entities).Error
}
