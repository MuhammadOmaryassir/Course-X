package repository

import "../entity"

type VideoRepository interface {
	Save(video *entity.Video) (*entity.Video, error)
	FindAll() ([]entity.Video, error)
}
