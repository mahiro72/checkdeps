package usecase

import (
	"a2/controller" // want "error: found bug in dependency import"
)

type A struct {
	con *controller.A
}
