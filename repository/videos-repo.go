package repository

import "cloudX/entity"

type VideoRepository interface {
	Save(post *entity.Video) (*entity.Video, error)
	FindAll() ([]entity.Video, error)
	FindByID(id string) (*entity.Video, error)
	Delete(video *entity.Video) error
}
