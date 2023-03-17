package controller

import (
	"a5/domain/repository" // want "error: found bug in dependency import"
)

type A struct {
	repo *repository.A
}
