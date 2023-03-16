package usecase

import (
	"a3/domain/repository"
	"a3/domain/model"
)

type A struct {
	repo repository.A
}

var _ *model.A = (*model.A)(nil)

