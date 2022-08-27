package usecase

import mongoRepo "github.com/zephyrzth/wdiet-be/repo/mongodb"

type usecase struct {
	mongoRepo mongoRepo.RepoInterface
}

func New(mongoRepo mongoRepo.RepoInterface) *usecase {
	return &usecase{
		mongoRepo: mongoRepo,
	}
}
