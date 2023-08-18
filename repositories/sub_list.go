package repositories

import (
	"moonlay/models"

	"gorm.io/gorm"
)

type SubListRepository interface {
	GetSubList(ID int) (models.SubList, error)
	GetAllSubLists() ([]models.SubList, error)
	GetPostImageSubs(ID int) ([]models.PostImageSub, error)
	GetAllPostImageSubs() ([]models.PostImageSub, error)
	CreateSubList(list models.SubList) (models.SubList, error)
	UpdateSubList(list models.SubList) (models.SubList, error)
	DeleteSubList(ID int) error
}

func RepositorySubList(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateSubList(sublist models.SubList) (models.SubList, error) {
	err := r.db.Preload("List").Create(&sublist).Error

	return sublist, err
}

func (r *repository) GetSubList(ID int) (models.SubList, error) {
	var list models.SubList
	err := r.db.First(&list, ID).Error

	return list, err
}

func (r *repository) GetPostImageSubs(ID int) ([]models.PostImageSub, error) {
	var postImageSub []models.PostImageSub
	err := r.db.Preload("PostImageSub").First(&postImageSub, ID).Error

	return postImageSub, err
}

func (r *repository) GetAllSubLists() ([]models.SubList, error) {
	var sublists []models.SubList
	err := r.db.Find(&sublists).Error

	return sublists, err
}

func (r *repository) GetAllPostImageSubs() ([]models.PostImageSub, error) {
	var postImageSubs []models.PostImageSub
	err := r.db.Find(&postImageSubs).Error

	return postImageSubs, err
}

func (r *repository) UpdateSubList(sublist models.SubList) (models.SubList, error) {
	err := r.db.Save(&sublist).Error
	return sublist, err
}

func (r *repository) DeleteSubList(ID int) error {
	return r.db.Delete(&models.SubList{}, ID).Error
}
