package service

import (
	"errors"
	"math/rand"
	"strconv"

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

func (*service) FindByID(id string) (*entity.Video, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}
