package usecase

import (
	"a4/domain/model"
	"a4/repository" // want "error: found bug in dependency import"
)

type A struct {
	repo repository.A
}

var _ *model.A = (*model.A)(nil)
