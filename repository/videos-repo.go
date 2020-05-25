package repository

import "CourseX/entity"

type VideoRepository interface {
	Save(video *entity.Video) (*entity.Video, error)
	FindAll() ([]entity.Video, error)
	FindByID(id string) (*entity.Video, error)
	Delete(video *entity.Video) error
}
