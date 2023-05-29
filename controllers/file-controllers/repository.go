package filecontrollers

import (
	"go-crud/models"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	CreateFile(input *models.FileModel) (*models.FileModel, int)

	GetAllFiles() ([]models.FileModel, int)

	DeleteFile(fileID uint) int
}

type repository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) CreateFile(input *models.FileModel) (*models.FileModel, int) {

	db := repo.db

	var file models.UserEntity

	checkIfFileExists := db.Select("*").Where("ID=?", input.ID).Find(&file)

	if checkIfFileExists.RowsAffected > 0 {
		return nil, http.StatusConflict
	}
	db.NewRecord(input)
	createFile := db.Create(&input)

	if createFile.Error != nil {
		return nil, http.StatusExpectationFailed
	}

	return input, http.StatusCreated
}

func (repo *repository) GetAllFiles() ([]models.FileModel, int) {

	return nil, 0
}

func (repo *repository) DeleteFile(fileID uint) int {

	return 0
}
