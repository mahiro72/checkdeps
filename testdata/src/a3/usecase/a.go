package usecase

import (
	"a3/domain/model"
	"a3/domain/repository"
)

type A struct {
	repo repository.A
}

var _ *model.A = (*model.A)(nil)
