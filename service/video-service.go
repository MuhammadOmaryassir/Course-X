package service

import (
	"errors"
	"math/rand"

	"../entity"
	"../repository"
)

type VideoService interface {
	Validate(video *entity.Video) error
	Create(video *entity.Video) (*entity.Video, error)
	FindAll() ([]entity.Video, error)
}

type service struct{}

var repo repository.VideoRepository

func NewVideoService(repository repository.VideoRepository) VideoService {
	repo = repository
	return &service{}
}

func (*service) Validate(video *entity.Video) error {
	if video == nil {
		err := errors.New("The video is empty")
		return err
	}
	if video.URL == "" {
		err := errors.New("The video title is empty")
		return err
	}
	return nil
}

func (*service) Create(video *entity.Video) (*entity.Video, error) {
	video.ID = rand.Int63()
	return repo.Save(video)
}

func (*service) FindAll() ([]entity.Video, error) {
	return repo.FindAll()
}
