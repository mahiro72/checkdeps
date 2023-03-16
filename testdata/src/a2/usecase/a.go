package usecase

import (
	"a/controller" // want "error: found bug in dependency import"
)

type A struct {
	con *controller.A
}
