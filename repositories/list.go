package repositories

import (
	"moonlay/models"

	"gorm.io/gorm"
)

type ListRepository interface {
	GetList(ID int) (models.List, error)
	GetAllLists() ([]models.List, error)
	CreateList(list models.List) (models.List, error)
	UpdateList(list models.List) (models.List, error)
	DeleteList(ID int) error
}

func RepositoryList(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateList(list models.List) (models.List, error) {
	err := r.db.Create(&list).Error

	return list, err
}

func (r *repository) GetList(ID int) (models.List, error) {
	var list models.List
	err := r.db.Preload("PostImage").Preload("SubList").First(&list, ID).Error

	return list, err
}

func (r *repository) GetAllLists() ([]models.List, error) {
	var lists []models.List
	err := r.db.Preload("PostImage").Preload("SubList").Find(&lists).Error

	return lists, err
}

func (r *repository) UpdateList(list models.List) (models.List, error) {
	err := r.db.Save(&list).Error
	return list, err
}

func (r *repository) DeleteList(ID int) error {
	return r.db.Delete(&models.List{}, ID).Error
}
