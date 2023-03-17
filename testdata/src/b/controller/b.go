package controller

import (
	usecase "b/repository" // want "error: found bug in dependency import"
)

type B struct {
	uc *usecase.B
}
